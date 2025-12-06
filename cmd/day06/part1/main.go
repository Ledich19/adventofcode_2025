package main

import (
	"adventofcode_2025/internal"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	Multiply = "*"
	Plus     = "+"
)

func extractActionAndNumbers(matrix [][]string, column int) (string, []int) {
	action := Plus
	numbers := make([]int, 0)

	for _, row := range matrix {
		cell := row[column]

		if cell != Plus && cell != Multiply {
			value, err := strconv.Atoi(cell)
			if err != nil {
				log.Fatalf("failed to parse %q: %v", cell, err)
			}
			numbers = append(numbers, value)
		}

		if cell == Plus || cell == Multiply {
			action = cell
			continue
		}
	}

	return action, numbers
}

func aggregate(action string, numbers []int) int {
	switch action {
	case Plus:
		sum := 0
		for _, n := range numbers {
			sum += n
		}
		return sum
	case Multiply:
		product := 1
		for _, n := range numbers {
			product *= n
		}
		return product
	default:
		return 0
	}
}

func main() {
	lines, err := internal.ReadLines("cmd/day06/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	result := 0
	matrix := make([][]string, 0)

	for _, line := range lines {
		parts := strings.Fields(line)
		matrix = append(matrix, parts)
	}

	for i := 0; i < len(matrix[0]); i++ {
		action, numbers := extractActionAndNumbers(matrix, i)
		result += aggregate(action, numbers)
	}

	fmt.Println("Day 06 Part 1:", result)
}
