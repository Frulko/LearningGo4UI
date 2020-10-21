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

func (r *rect) setSize(w int, h int) {
	r.width = w
	r.height = h
}

func (r *rect) setPosition(x int, y int) {
	r.x = x
	r.y = y
}

func (r *rect) getSize() []int {

	var i []int
	i = append(i, r.width)
	i = append(i, r.height)

	return i
}
