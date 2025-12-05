package main

import (
	"adventofcode_2025/internal"
	"adventofcode_2025/internal/numbers"
	"adventofcode_2025/internal/ranges"
	"fmt"
	"log"
	"slices"
)

func main() {
	countFresh := 0

	lines, err := internal.ReadLines("cmd/day05/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	index := slices.Index(lines, "")
	fmt.Println("Index of empty line:", index)

	parsedRanges := ranges.ParseMany(lines[:index])
	numbers := numbers.ParseMany(lines[index+1:])

	for _, n := range numbers {
		for _, r := range parsedRanges {
			if n >= r.Start && n <= r.End {
				countFresh++
				break
			}
		}
	}

	fmt.Println("Day 05 Part 1:", countFresh)
}
