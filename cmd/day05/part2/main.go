package main

import (
	"adventofcode_2025/internal"
	"adventofcode_2025/internal/ranges"
	"fmt"
	"log"
	"slices"
	"sort"
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

	sort.Slice(parsedRanges, func(i, j int) bool {
		return parsedRanges[i].Start < parsedRanges[j].Start
	})

	crossRanges := ranges.MergeOverlapping(parsedRanges)

	for _, r := range crossRanges {
		countFresh += r.End - r.Start + 1
	}

	fmt.Println("Day 05 Part 1:", countFresh)
}
