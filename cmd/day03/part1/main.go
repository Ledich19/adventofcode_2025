package main

import (
	"adventofcode_2025/internal"
	"adventofcode_2025/internal/digits"
	"adventofcode_2025/internal/math"
	"fmt"
	"log"
	"strconv"
)

func main() {
	resultSum := 0

	text, err := internal.ReadLines("cmd/day03/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, line := range text {
		step := 12

		numbers, err := digits.LineToDigits(line)
		if err != nil {
			log.Fatalf("failed to parse digits: %v", err)
		}
		maxIndex := math.MaxIndex(numbers[:len(numbers)-step])
		secondRange := numbers[maxIndex+1:]
		maxSecondIndex := math.MaxIndex(secondRange)

		concat := fmt.Sprintf("%d%d", numbers[maxIndex], secondRange[maxSecondIndex])
		value, err := strconv.Atoi(concat)
		if err != nil {
			log.Fatalf("failed to parse %s: %v", concat, err)
		}

		resultSum += value
		fmt.Println(numbers, "max:", numbers[maxIndex], "max second:", secondRange[maxSecondIndex])
	}

	fmt.Println("Final value:", resultSum)
}
