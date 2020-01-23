package strategy

import (
	"github.com/tinychameleon/blocky/grid"
	"github.com/tinychameleon/blocky/shape"
	"image/color"
)

func Rects(g grid.Grid) []shape.Interface {
	var shapes []shape.Interface
	width, height := g.Bounds()
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			c := g.At(x, y)
			if c == nil {
				continue
			}

			var r shape.Interface
			l := left(g, x, y, *c)
			d := down(g, x, y, *c)
			if l >= d {
				r = shape.Rect(x, y, l, 1, *c)
				ub := x + l
				for x2 := x; x2 < ub; x2++ {
					g.EraseAt(x2, y)
				}
			} else {
				r = shape.Rect(x, y, 1, d, *c)
				ub := y + d
				for y2 := y; y2 < ub; y2++ {
					g.EraseAt(x, y2)
				}
			}

			shapes = append(shapes, r)
		}
	}
	return shapes
}

func left(g grid.Grid, x int, y int, c color.RGBA) int {
	width, _ := g.Bounds()
	count := 1
	for i := x + 1; i < width; i++ {
		cn := g.At(i, y)
		if cn == nil || *cn != c {
			break
		}
		count++
	}
	return count
}

func down(g grid.Grid, x int, y int, c color.RGBA) int {
	_, height := g.Bounds()
	count := 1
	for i := y + 1; i < height; i++ {
		cn := g.At(x, i)
		if cn == nil || *cn != c {
			break
		}
		count++
	}
	return count
}
