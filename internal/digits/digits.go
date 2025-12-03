package digits

import (
	"strconv"
	"strings"
)

func LineToDigits(line string) ([]int, error) {
	parts := strings.Split(line, "")
	digits := make([]int, len(parts))

	for i, s := range parts {
		n, err := strconv.Atoi(s)
		if err != nil {
			return nil, err
		}
		digits[i] = n
	}

	return digits, nil
}
