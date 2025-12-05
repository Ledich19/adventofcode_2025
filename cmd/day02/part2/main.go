package main

import (
	"adventofcode_2025/internal"
	"adventofcode_2025/internal/numbers"
	"adventofcode_2025/internal/ranges"
	"fmt"
	"log"
	"strconv"
	"strings"
)

func compareSlices(i int, numberStr string, k int, sliceFist string) bool {
	length := numbers.Length(i)

	for j := 1; j <= length/k; j++ {
		if k*j+k > length {
			break
		}
		nextSlice := numberStr[k*j : k*j+k]

		if sliceFist != nextSlice {
			return false
		}
	}

	return length%k == 0
}

func hasRepeatingSlice(i int, numberStr string, maxLength int) bool {
	for k := 1; k <= maxLength; k++ {
		sliceFist := numberStr[:k]
		if compareSlices(i, numberStr, k, sliceFist) {
			return true
		}
	}
	return false
}

func main() {
	sum := 0

	text, err := internal.ReadText("cmd/day02/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	parts := strings.Split(text, ",")
	parsedRanges := ranges.ParseMany(parts)

	for _, r := range parsedRanges {
		for i := r.Start; i <= r.End; i++ {
			maxLength := numbers.Length(i) / 2
			numberStr := strconv.Itoa(i)

			if hasRepeatingSlice(i, numberStr, maxLength) {
				sum += i
			}
		}
	}

	fmt.Println("Final value:", sum)

}
