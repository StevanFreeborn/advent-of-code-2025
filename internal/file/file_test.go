package file_test

import (
	"slices"
	"testing"

	"github.com/StevanFreeborn/advent-of-code-2025/internal/file"
)

func TestAllReadLines(t *testing.T) {
	expectedLines := []string{
		"hello",
		"world",
	}

	result := file.ReadAllLines("test.txt")

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

func TestReadLines(t *testing.T) {
	t.Run("it should return an empty line if the file does not exist", func(t *testing.T) {
		expectedLines := []string{""}
		lines := []string{}

		for line := range file.ReadLines("non-existent-file.txt") {
			lines = append(lines, line)
		}

		if slices.Equal(lines, expectedLines) == false {
			t.Errorf("got %s but expected %s", lines, expectedLines)
		}
	})

	t.Run("it should read all lines from the file", func(t *testing.T) {
		expectedLines := []string{
			"hello",
			"world",
		}

		lines := []string{}

		for line := range file.ReadLines("test.txt") {
			lines = append(lines, line)
		}

		if slices.Equal(lines, expectedLines) == false {
			t.Errorf("got %s but expected %s", lines, expectedLines)
		}
	})

	t.Run("it should stop reading lines when iteration is stopped", func(t *testing.T) {
		expectedLines := []string{
			"hello",
		}

		lines := []string{}

		for line := range file.ReadLines("test.txt") {
			if len(lines) >= 1 {
				break
			}

			lines = append(lines, line)
		}

		if slices.Equal(lines, expectedLines) == false {
			t.Errorf("got %s but expected %s", lines, expectedLines)
		}
	})
}
