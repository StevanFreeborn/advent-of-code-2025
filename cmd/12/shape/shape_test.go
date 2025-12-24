package shape_test

import (
	"strings"
	"testing"

	"github.com/StevanFreeborn/advent-of-code-2025/cmd/12/shape"
)

func TestGenerateVariants(t *testing.T) {
	t.Run("works for shape 0", func(t *testing.T) {
		expectedVariants := [][]string{
			{
				"###",
				"##.",
				"##.",
			},
			{
				"###",
				".##",
				".##",
			},
			{
				"###",
				"###",
				"#..",
			},
			{
				"###",
				"###",
				"..#",
			},
			{
				".##",
				".##",
				"###",
			},
			{
				"##.",
				"##.",
				"###",
			},
		}

		lines := []string{
			"0:",
			"###",
			"##.",
			"##.",
		}

		shape := shape.From(lines)

		result := shape.GenerateVariants()

		for _, ev := range expectedVariants {
			expected := strings.Join(ev, "\n")

			found := false

			for _, v := range result {
				if strings.TrimSpace(v.String()) == expected {
					found = true
				}
			}

			if found == false {
				t.Errorf("expected to find variant but did not:\n%v", expected)
			}
		}
	})

	t.Run("works for shape 1", func(t *testing.T) {
		expectedVariants := [][]string{
			{
				"###",
				"##.",
				".##",
			},
			{
				".##",
				"###",
				"#.#",
			},
			{
				"##.",
				".##",
				"###",
			},
			{
				"#.#",
				"###",
				"##.",
			},
			{
				"###",
				".##",
				"##.",
			},
			{
				"##.",
				"###",
				"#.#",
			},
			{
				"##.",
				"###",
				"#.#",
			},
			{
				".##",
				"##.",
				"###",
			},
			{
				"#.#",
				"###",
				".##",
			},
		}

		lines := []string{
			"0:",
			"###",
			"##.",
			".##",
		}

		shape := shape.From(lines)

		result := shape.GenerateVariants()

		for _, ev := range expectedVariants {
			expected := strings.Join(ev, "\n")

			found := false

			for _, v := range result {
				if strings.TrimSpace(v.String()) == expected {
					found = true
				}
			}

			if found == false {
				t.Errorf("expected to find variant but did not:\n%v", expected)
			}
		}
	})
}
