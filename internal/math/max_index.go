package math

func MaxIndex(arr []int) int {
	if len(arr) == 0 {
		return -1
	}

	maxIdx := 0
	maxVal := arr[0]

	for i := 1; i < len(arr); i++ {
		if arr[i] > maxVal {
			maxVal = arr[i]
			maxIdx = i
		}
	}

	return maxIdx
}
