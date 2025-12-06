package array

func Every[T any](items []T, predicate func(T) bool) bool {
	for _, v := range items {
		if !predicate(v) {
			return false
		}
	}
	return true
}
