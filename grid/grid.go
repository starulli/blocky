// Package grid provides a 2D colour array with associated metadata.
package grid

import (
	"image"
	"image/color"
)

// Grid is a 2D colour array populated from image colour data.
type Grid interface {
	At(int, int) color.RGBA
	Bounds() (int, int)
}

// New creates a Grid from the given image colour data. It stores all data as
// color.RGBA values.
func New(m image.Image) Grid {
	b := m.Bounds()
	c := make([]*color.RGBA, b.Max.X * b.Max.Y)
	for y := 0; y < b.Max.Y; y++ {
		for x := 0; x < b.Max.X; x++ {
			i := y * b.Max.X + x
			rgba := color.RGBAModel.Convert(m.At(x, y)).(color.RGBA)
			c[i] = &rgba
		}
	}
	return gridarray{width: b.Max.X, height: b.Max.Y, colours: c}
}

type gridarray struct {
	width, height int
	colours []*color.RGBA
}

func (g gridarray) At(x int, y int) color.RGBA {
	return *g.colours[g.width*y+x]
}

func (g gridarray) Bounds() (int, int) {
	return g.width, g.height
}
