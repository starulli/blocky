package strategy

import (
	"github.com/tinychameleon/blocky/grid"
	"github.com/tinychameleon/blocky/shape"
)

// Pixels transforms a Grid into individual pixel shapes.
func Pixels(g grid.Grid) []shape.Interface {
	width, height := g.Bounds()
	shapes := make([]shape.Interface, width * height)
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			i := y * width + x
			shapes[i] = shape.Pixel(x, y, g.At(x, y))
		}
	}
	return shapes
}
