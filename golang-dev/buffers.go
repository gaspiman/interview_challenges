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
	"strconv"
	"strings"
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
	for _, file := range files {
		file_path := filepath.Join(path, file)
		inputFile, err := os.Open(file_path)
		checkErr(err)
		scanner := bufio.NewScanner(inputFile) // Intializing a buffered scanner
		for scanner.Scan() {
			words := strings.Fields(scanner.Text()) // reading line by line
			for _, word := range words {
				word = strings.ToLower(word)
				if _, ok := m[word]; ok {
					m[word] += 1
					continue
				}
				m[word] = 1
			}
		}
		/*
			Even better and more resilient code would be:
			scanner := bufio.NewScanner(gr)
			// Setting buffer size in case line is bigger than the available RAM
			buf := make([]byte, 0, 10*1024*1024) // 10MB buffer
			scanner.Buffer(buf, 10*1024*1024) // 10MB buffer
			for scanner.Scan() {
				sanner.Text()
				// word counting logic
			}
		*/

		checkErr(scanner.Err())
	}
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

// Function that writes the word count in the CSV file
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
