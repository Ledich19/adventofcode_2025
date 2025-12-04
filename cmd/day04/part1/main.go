package main

import (
	"adventofcode_2025/internal"
	"adventofcode_2025/internal/grid"
	"fmt"
	"log"
)

const (
	Roll  rune = '@'
	Empty rune = '.'
)

func countRolls(matrix [][]rune, points []grid.Point) int {
	count := 0
	for _, p := range points {
		if p.Y >= 0 && p.Y < len(matrix) && p.X >= 0 && p.X < len(matrix[p.Y]) {
			if matrix[p.Y][p.X] == Roll {
				count++
			}
		}
	}
	return count
}

func main() {
	matrix, err := internal.ReadMatrix("cmd/day04/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	result := 0

	grid.ForEachCell(matrix, func(x, y int, cell rune) {
		mask := grid.GetMask(x, y)
		if cell == Roll && countRolls(matrix, mask) < 4 {
			result += 1
		}
	})

	fmt.Println("Day 04 Part 1:", result)
}
