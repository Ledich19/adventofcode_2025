package main

import (
	"adventofcode_2025/internal"
	"adventofcode_2025/internal/ranges"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func hasEqualHalves(n int) bool {
	numberStr := strconv.Itoa(n)
	if len(numberStr)%2 != 0 {
		return false
	}

	mid := len(numberStr) / 2
	return numberStr[:mid] == numberStr[mid:]
}

func main() {
	countInvalid := 0

	text, err := internal.ReadText("cmd/day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(text, ",")
	parsedRanges := ranges.ParseMany(parts)

	for _, r := range parsedRanges {
		for i := r.Start; i <= r.End; i++ {
			if hasEqualHalves(i) {
				countInvalid += i
			}
		}
	}

	fmt.Println("Final value:", countInvalid)
}
