package main

import (
	"adventofcode_2025/internal"
	"fmt"
	"strconv"
)

func main() {
	START_POSITION := 50
	position := START_POSITION

	lines, _ := internal.ReadLines("cmd/day01/input.txt")

	countNuls := 0

	for _, line := range lines {
		first := line[0]
		numberStr := line[1:]
		number, _ := strconv.Atoi(numberStr)

		switch first {
		case 'L':
			first0 := position
			if first0 == number {
				countNuls += 1
			}
			if first0 < number && first0 != 0 {
				countNuls += 1
			}
			if first0 < number {
				countNuls += (number - first0) / 100
			}
			position = (position - number%100 + 100) % 100

		case 'R':
			first0 := 100 - position

			if first0 == number {
				countNuls += 1
			}
			if first0 < number && first0 != 0 {
				countNuls += 1
			}
			if first0 < number {
				countNuls += (number - first0) / 100
			}
			position = (position + number) % 100
		}

	}

	fmt.Println("Count of nul positions:", countNuls)
}
