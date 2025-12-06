package grid

type Point struct {
	X int
	Y int
}

// GetMask returns coordinates of 8 neighboring cells around (x, y).
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
