package ranges

import (
	"strconv"
	"strings"
)

type Range struct {
	Start int
	End   int
}

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

func MergeOverlapping(sorted []Range) []Range {
	if len(sorted) == 0 {
		return nil
	}

	merged := []Range{sorted[0]}

	for _, r := range sorted[1:] {
		last := &merged[len(merged)-1]
		if r.Start <= last.End {
			if r.End > last.End {
				last.End = r.End
			}
		} else {
			merged = append(merged, r)
		}
	}

	return merged
}
