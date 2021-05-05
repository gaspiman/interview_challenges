/*
Given:
- Input folder containing text files
- Count the number of time each word is repeated
- Export the results in a CSV file (ex: word, 10)
*/

package main

import (
	"encoding/csv"
	"io/ioutil"
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
		data, err := ioutil.ReadFile(file_path) // The entire file is read in memery
		checkErr(err)
		words := strings.Fields(string(data))
		for _, word := range words {
			word = strings.ToLower(word)
			if _, ok := m[word]; ok {
				m[word] += 1
				continue
			}
			m[word] = 1
		}
	}
	outputFile, err := os.Create("result.csv")
	checkErr(err)
	defer outputFile.Close()
	writer := csv.NewWriter(outputFile)
	defer writer.Flush()
	for word, count := range m {
		writer.Write([]string{word, strconv.Itoa(count)})
	}
}

// Function to extract all files from a folder
func get_files(path string) []string {
	files := []string{}
	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		files = append(files, path)
		return nil
	})
	checkErr(err)
	return files
}
