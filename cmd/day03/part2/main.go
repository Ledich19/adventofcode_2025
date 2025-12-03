package main

import (
	"adventofcode_2025/internal"
	"adventofcode_2025/internal/digits"
	"adventofcode_2025/internal/math"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func selectMaxSequence(numbers []int, length int) []int {
	result := make([]int, 0, length)

	start := 0
	n := len(numbers)
	for i := 0; i < length; i++ {
		tail := length - (i + 1)
		end := n - tail

		idx := math.MaxIndex(numbers[start:end])
		value := numbers[start+idx]
		result = append(result, value)
		start = start + idx + 1
	}

	return result
}

func main() {
	NUMBER_LENGTH := 12
	resultSum := 0

	text, err := internal.ReadLines("cmd/day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range text {
		numbers, _ := digits.LineToDigits(line)
		result := selectMaxSequence(numbers, NUMBER_LENGTH)
		numberStr := strings.Builder{}

		for _, digit := range result {
			numberStr.WriteString(strconv.Itoa(digit))
		}

		number, _ := strconv.Atoi(numberStr.String())
		resultSum += number
	}

	fmt.Println("Final value:", resultSum)
}
