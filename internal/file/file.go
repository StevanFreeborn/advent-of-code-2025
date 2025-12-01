// Package file provides useful methods for interacting with files
package file

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(filePath string) []string {
	lines := []string{}

	file, openErr := os.Open(filePath)

	if openErr != nil {
		return lines
	}

	scanner := bufio.NewScanner(file)

	for hasLine := scanner.Scan(); hasLine; hasLine = scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		lines = append(lines, line)
	}

	return lines
}
