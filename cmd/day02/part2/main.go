package main

import (
	"adventofcode_2025/internal"
	"adventofcode_2025/internal/numbers"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func parseRanges(text string) []Range {
	text = strings.TrimSpace(text)
	parts := strings.Split(text, ",")
	result := make([]Range, 0, len(parts))

	for _, p := range parts {
		edges := strings.Split(p, "-")
		start, _ := strconv.Atoi(edges[0])
		end, _ := strconv.Atoi(edges[1])

		result = append(result, Range{
			start: start,
			end:   end,
		})
	}

	return result
}

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

	ranges := parseRanges(text)

	for _, r := range ranges {
		for i := r.start; i <= r.end; i++ {
			maxLength := numbers.Length(i) / 2
			numberStr := strconv.Itoa(i)

			if hasRepeatingSlice(i, numberStr, maxLength) {
				sum += i
			}
		}
	}

	fmt.Println("Final value:", sum)

}
