package ui

// extend yoga view and properties / methods
type rect struct {
	width  int
	height int
	x      int
	y      int
}

func Rect(w int, h int, x int, y int) *rect {
	r := rect{width: w, height: h, x: x, y: y}
	return &r
}

func (r rect) GetWidth() int {
	return r.width
}
