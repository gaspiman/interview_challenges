/*
Given:
- Input folder containing text files
- Count the number of time each word is repeated
- Export the results in a CSV file (ex: word, 10)
*/
package main

import (
	"bufio"
	"encoding/csv"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
	"sync"
)

var path = "/path/to/folder"
var output = "/path/to/output/file.csv"

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	m := map[string]int{}
	files := get_files(path)
	inCH := make(chan string) // Creating the input communication channel
	outCH := make(chan map[string]int)
	wg := new(sync.WaitGroup) // Wait group

	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go func(inCH chan string, outCH chan map[string]int, wg *sync.WaitGroup) {
			localMap := map[string]int{} // Each go routine has its own localMap
			for file_path := range inCH {
				inputFile, err := os.Open(file_path)
				checkErr(err)
				scanner := bufio.NewScanner(inputFile)
				// Setting buffer size in case line is bigger than the available RAM
				buf := make([]byte, 0, 10*1024*1024) // 10MB buffer
				scanner.Buffer(buf, 10*1024*1024)    // 10MB buffer
				for scanner.Scan() {
					words := strings.Fields(scanner.Text()) // reading lin by line
					for _, word := range words {
						word = strings.ToLower(word)
						if _, ok := m[word]; ok {
							localMap[word] += 1
							continue
						}
						localMap[word] = 1
					}
				}
			}
			// We have accumulated the word counts - let's send them to the main map
			outCH <- localMap
		}(inCH, outCH, wg)
	}
	// The Go function listens to the Output communication channel and records the word counts in the main map
	go func(outCH chan map[string]int) {
		// The on the output channel we receive the local maps from the go routines
		for localMap := range outCH {
			for word, count := range localMap {
				if _, ok := m[word]; ok {
					m[word] += count
					continue
				}
				m[word] = 1
			}
		}
	}(outCH)

	for _, file := range files {
		file_path := filepath.Join(path, file)
		inCH <- file_path
	}
	close(inCH) // Closing the input communication channel
	wg.Wait()   // Wating for all goroutines to finish

	// Write the results to the CSV file
	writeToFile(output, m)
}

func get_files(path string) []string {
	files := []string{}
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	checkErr(err)
	return files
}

func writeToFile(path string, m map[string]int) {
	outputFile, err := os.Create(output)
	checkErr(err)
	defer outputFile.Close()
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()
	for word, count := range m {
		writer.Write([]string{word, strconv.Itoa(count)})
	}
}
