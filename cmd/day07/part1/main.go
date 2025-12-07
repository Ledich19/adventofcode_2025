package main

import (
	"adventofcode_2025/internal"
	"adventofcode_2025/internal/grid"
	"fmt"
	"log"
)

const (
	Start   = 'S'
	Empty   = '.'
	Divider = '^'
	Beam    = '|'
)

func main() {
	matrix, err := internal.ReadMatrix("cmd/day07/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	countDividerActions := 0

	grid.ForEachCell(matrix, func(x, y int, value rune) {
		isTopBeam := y-1 >= 0 && matrix[y-1][x] == Beam
		isBottomLeftEmpty := x-1 >= 0 && y+1 < len(matrix) && matrix[y+1][x-1] == Empty
		isBottomRightEmpty := x+1 < len(matrix[y]) && y+1 < len(matrix) && matrix[y+1][x+1] == Empty

		if value == Start && y+1 < len(matrix) {
			matrix[y+1][x] = Beam
		}

		if isTopBeam && value == Empty {
			matrix[y][x] = Beam
		}

		if value == Divider && y < len(matrix)-1 {
			wasDividerUsed := false

			if isBottomLeftEmpty && isTopBeam {
				matrix[y][x-1] = Beam
				matrix[y+1][x-1] = Beam
				wasDividerUsed = true
			}

			if isBottomRightEmpty && isTopBeam {
				matrix[y][x+1] = Beam
				matrix[y+1][x+1] = Beam
				wasDividerUsed = true
			}

			if wasDividerUsed {
				countDividerActions++
			}

		}
	})

	fmt.Println("Day 07 Part 1:", countDividerActions)
}
