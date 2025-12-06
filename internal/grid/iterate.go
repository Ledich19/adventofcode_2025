package grid

// ForEachCell iterates over all cells of a rune matrix.
func ForEachCell(matrix [][]rune, fn func(x, y int, value rune)) {
	for y, row := range matrix {
		for x, value := range row {
			fn(x, y, value)
		}
	}
}
