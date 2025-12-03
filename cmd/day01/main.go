package main

import (
	"adventofcode_2025/internal"
	"fmt"
	"log"
	"strconv"
)

func main() {
	START_POSITION := 50

	lines, err := internal.ReadLines("cmd/day01/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	value := START_POSITION
	countNuls := 0

	for _, line := range lines {
		first := line[0]
		numberStr := line[1:]
		number, _ := strconv.Atoi(numberStr)

		switch first {
		case 'L':
			value -= number
		case 'R':
			value += number
		}

		if value == 0 || value%100 == 0 {
			countNuls += 1
		}

	}
	fmt.Println("Final value:", value)
	fmt.Println("Count of nul positions:", countNuls)
}
