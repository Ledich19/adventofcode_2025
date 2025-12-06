package ranges

import (
	"strconv"
	"strings"
)

func ParseMany(parts []string) []Range {
	result := make([]Range, 0, len(parts))

	for _, p := range parts {
		edges := strings.Split(p, "-")
		start, _ := strconv.Atoi(edges[0])
		end, _ := strconv.Atoi(edges[1])

		result = append(result, Range{
			Start: start,
			End:   end,
		})
	}

	return result
}
