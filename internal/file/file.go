// Package file provides useful methods for interacting with files
package file

import (
	"bufio"
	"os"
)

// TODO: Implement a StreamLines method
// in the file package

func ReadLines(filePath string) []string {
	lines := []string{}

	file, openErr := os.Open(filePath)

	if openErr != nil {
		return lines
	}

	scanner := bufio.NewScanner(file)

	for hasLine := scanner.Scan(); hasLine; hasLine = scanner.Scan() {
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
