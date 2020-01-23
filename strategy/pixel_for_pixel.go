package strategy

import (
	"github.com/tinychameleon/blocky/shape"
	"image"
)

func PixelForPixel(m image.Image) []shape.Interface {
	var sh []shape.Interface
	for y := 0; y < m.Bounds().Max.Y; y++ {
		for x := 0; x < m.Bounds().Max.X; x++ {
			sh = append(sh, shape.Pixel(x, y, m.At(x, y)))
		}
	}
	return sh
}
