package ranges

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
