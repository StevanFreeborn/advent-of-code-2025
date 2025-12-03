package file_test

import (
	"slices"
	"testing"

	"github.com/StevanFreeborn/advent-of-code-2025/internal/file"
)

func TestReadLines(t *testing.T) {
	expectedLines := []string{
		"hello",
		"world",
	}

	result := file.ReadLines("test.txt")

	if slices.Equal(result, expectedLines) == false {
		t.Errorf("got %s but expected %s", result, expectedLines)
	}
}

func TestReadAllText(t *testing.T) {
	expectedText := "hello\r\nworld\r\n"

	result := file.ReadAllText("test.txt")

	if result != expectedText {
		t.Errorf("got %q but expected %q", result, expectedText)
	}
}
