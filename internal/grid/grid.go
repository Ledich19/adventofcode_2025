package grid

type Point struct {
	X int
	Y int
}

func GetMask(x, y int) []Point {
	return []Point{
		{X: x - 1, Y: y - 1},
		{X: x, Y: y - 1},
		{X: x + 1, Y: y - 1},
		{X: x - 1, Y: y},
		{X: x + 1, Y: y},
		{X: x - 1, Y: y + 1},
		{X: x, Y: y + 1},
		{X: x + 1, Y: y + 1},
	}
}

// ForEachCell walks a rune matrix and calls fn for each cell.
func ForEachCell(matrix [][]rune, fn func(x, y int, value rune)) {
	for y, row := range matrix {
		for x, value := range row {
			fn(x, y, value)
		}
	}
}
