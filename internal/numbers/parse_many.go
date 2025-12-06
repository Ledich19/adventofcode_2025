package numbers

import "strconv"

func ParseMany(parts []string) []int {
	result := make([]int, 0, len(parts))

	for _, p := range parts {
		value, _ := strconv.Atoi(p)
		result = append(result, value)
	}

	return result
}
