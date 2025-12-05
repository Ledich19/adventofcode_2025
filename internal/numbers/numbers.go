package numbers

import "strconv"

func Length(n int) int {
	if n == 0 {
		return 1
	}

	if n < 0 {
		n = -n
	}

	length := 0
	for n > 0 {
		n /= 10
		length++
	}

	return length
}

func ParseMany(parts []string) []int {
	result := make([]int, 0, len(parts))

	for _, p := range parts {
		value, _ := strconv.Atoi(p)
		result = append(result, value)
	}

	return result
}
