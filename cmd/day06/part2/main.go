package main

import (
	"adventofcode_2025/internal"
	"adventofcode_2025/internal/array"
	"fmt"
	"log"
	"strconv"
	"strings"
)

const (
	Multiply = "*"
	Plus     = "+"
)

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
		parts := strings.Split(line, "")
		matrix = append(matrix, array.Reverse(parts))
	}

	numbers := make([]int, 0)
	action := Plus

	for i := 0; i < len(matrix[0]); i++ {
		group := make([]string, 0)

		for _, row := range matrix {
			cell := row[i]
			if cell == Plus || cell == Multiply {
				action = cell
			} else {
				group = append(group, cell)
			}
		}

		groupStr := strings.TrimSpace(strings.Join(group, ""))

		if groupStr != "" {
			number, _ := strconv.Atoi(groupStr)
			numbers = append(numbers, number)
		}

		if groupStr == "" || len(matrix[0])-1 == i {
			result += aggregate(action, numbers)
			numbers = make([]int, 0)
		}
	}

	fmt.Println("Day 06 Part 1:", result)
}

// 10388762421009
