package array

func Reverse[T any](items []T) []T {
	result := make([]T, len(items))
	for i, v := range items {
		result[len(items)-1-i] = v
	}
	return result
}
