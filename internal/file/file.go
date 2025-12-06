// Package file provides useful methods for interacting with files
package file

import (
	"bufio"
	"iter"
	"os"
)

func ReadLines(filePath string) iter.Seq[string] {
	return func(yield func(string) bool) {
		file, err := os.Open(filePath)

		if err != nil {
			yield("")
			return
		}

		defer file.Close()

		scanner := bufio.NewScanner(file)

		for scanner.Scan() {
			if !yield(scanner.Text()) {
				return
			}
		}
	}
}

func ReadAllLines(filePath string) []string {
	lines := []string{}

	file, openErr := os.Open(filePath)

	if openErr != nil {
		return lines
	}

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}

func ReadAllText(filePath string) string {
	fileContent, readErr := os.ReadFile(filePath)

	if readErr != nil {
		return ""
	}

	return string(fileContent)
}
