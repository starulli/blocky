package strategy

import (
	"github.com/tinychameleon/blocky/grid"
	"github.com/tinychameleon/blocky/shape"
)

// Pixels transforms a Grid into individual pixel shapes.
func Pixels(g grid.Grid) []shape.Interface {
	width, height := g.Bounds()
	var shapes []shape.Interface
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c := g.At(x, y)
			if c == nil {
				continue
			}
			shapes = append(shapes, shape.Pixel(x, y, *c))
		}
	}
	return shapes
}
