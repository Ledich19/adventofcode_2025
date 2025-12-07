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

	countPath := 0

	grid.ForEachCell(matrix, func(x, y int, value rune) {
		isTopBeam := y-1 >= 0 && matrix[y-1][x] == Beam
		isBottomLeftEmpty := x-1 >= 0 && y+1 < len(matrix) && matrix[y+1][x-1] == Empty
		isBottomRightEmpty := x+1 < len(matrix[y]) && y+1 < len(matrix) && matrix[y+1][x+1] == Empty

		if value == Start {
			matrix[y+1][x] = Beam
		}

		if isTopBeam && value == Empty {
			matrix[y][x] = Beam
		}

		if value == Divider {

			if isBottomLeftEmpty && isTopBeam {
				matrix[y][x-1] = Beam
				matrix[y+1][x-1] = Beam
			}

			if isBottomRightEmpty && isTopBeam {
				matrix[y][x+1] = Beam
				matrix[y+1][x+1] = Beam
			}

		}
	})

	countMatrix := make([][]int, len(matrix))
	for y := range matrix {
		countMatrix[y] = make([]int, len(matrix[y]))
	}

	grid.ForEachCell(matrix, func(x, y int, value rune) {
		isBeamBefore := y-1 >= 0 && matrix[y-1][x] == Beam

		if value == Beam && y > 0 && matrix[y-1][x] == Start {
			countMatrix[y][x] += 1
		}

		if isBeamBefore {
			countMatrix[y][x] += countMatrix[y-1][x]
		}

		if x-1 >= 0 && value == Divider && isBeamBefore {
			countMatrix[y][x-1] += countMatrix[y-1][x]
		}

		if x+1 < len(matrix[y]) && value == Divider && isBeamBefore {
			countMatrix[y][x+1] += countMatrix[y-1][x]
		}

	})

	for _, value := range countMatrix[len(countMatrix)-1] {
		if value > 0 {
			countPath += value
		}
	}

	fmt.Println("Day 07 Part 2:", countPath)
}
