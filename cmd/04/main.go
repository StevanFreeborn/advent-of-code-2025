package main

import (
	"github.com/StevanFreeborn/advent-of-code-2025/internal/file"
)

const PaperRollCharacter = "@"
const EmptyCharacter = "."

type move struct {
	NumberOfRows    int
	NumberOfColumns int
}

var (
	up        = move{NumberOfRows: -1, NumberOfColumns: 0}
	down      = move{NumberOfRows: 1, NumberOfColumns: 0}
	right     = move{NumberOfRows: 0, NumberOfColumns: 1}
	left      = move{NumberOfRows: 0, NumberOfColumns: -1}
	upRight   = move{NumberOfRows: -1, NumberOfColumns: 1}
	upLeft    = move{NumberOfRows: -1, NumberOfColumns: -1}
	downRight = move{NumberOfRows: 1, NumberOfColumns: 1}
	downLeft  = move{NumberOfRows: 1, NumberOfColumns: -1}
)

var moves = []move{
	up,
	down,
	right,
	left,
	upRight,
	upLeft,
	downRight,
	downLeft,
}

func SolvePartOne(filePath string) int {
	input := file.ReadAllLines(filePath)
	numberOfRows := len(input)
	numberOfColumns := len(input[0])

	total := 0

	for row := range numberOfRows {
		for column := range numberOfColumns {
			value := string(input[row][column])

			if value != PaperRollCharacter {
				continue
			}

			numberOfPaperRolls := 0

			for _, move := range moves {
				neighborRow := row + move.NumberOfRows
				neighborColumn := column + move.NumberOfColumns

				if neighborRow > numberOfRows-1 || neighborRow < 0 {
					continue
				}

				if neighborColumn > numberOfColumns-1 || neighborColumn < 0 {
					continue
				}

				neighborValue := string(input[neighborRow][neighborColumn])

				if neighborValue != PaperRollCharacter {
					continue
				}

				numberOfPaperRolls++
			}

			if numberOfPaperRolls < 4 {
				total++
			}
		}
	}

	return total
}
