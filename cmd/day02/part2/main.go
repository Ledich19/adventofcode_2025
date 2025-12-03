package main

import (
	"adventofcode_2025/internal"
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

	ranges := parseRanges(text)

	for _, r := range ranges {
		for i := r.start; i <= r.end; i++ {
			if hasEqualHalves(i) {
				countInvalid += i
			}
		}
	}

	fmt.Println("Final value:", countInvalid)
}
