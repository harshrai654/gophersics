package ch1

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func Dup2() {
	files := os.Args[1:]
	freq := make(map[string]map[string]int)

	for _, fileName := range files {
		absPath := ""
		var err error
		if filepath.IsAbs(fileName) {
			absPath = fileName
		} else {
			absPath, err = filepath.Abs(fileName)
			if err != nil {
				fmt.Fprintf(os.Stderr, "%v", err)
				return
			}
			
		}
		
		file, err := os.Open(absPath)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%v", err)
			return
		}

		countLines(file, freq)
	}

	fmt.Println("Results...")

	for line, fileMap := range freq {
		fileNamesString := ""
		totalCount := 0

		for fileName, count := range fileMap {
			fileNamesString += fileName + ","
			totalCount += count
		}

		fmt.Printf("Line: %s | Files: %s | Total count: %d\n", line, fileNamesString, totalCount);
	}
}

func countLines(file *os.File, freq map[string]map[string]int) {
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		if freq[scanner.Text()] == nil {
			freq[scanner.Text()] = make(map[string]int)
		} 
		freq[scanner.Text()][file.Name()]++
	}
}