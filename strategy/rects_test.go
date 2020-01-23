package strategy

import (
	"github.com/tinychameleon/blocky/grid"
	"github.com/tinychameleon/blocky/shape"
	"github.com/tinychameleon/blocky/testdata"
	"image/color"
	"testing"
)

func TestRects(t *testing.T) {
	g := grid.New(testdata.Image())
	g.EraseMatching(func(c color.RGBA) bool {
		return c.A == 0x00
	})

	got := Rects(g)
	if len(got) != len(rects) {
		t.Errorf("Got len %v, want %v", len(got), len(rects))
	}
	for i, r := range rects {
		if got[i] != r {
			t.Errorf("Index %v: got %#v, want %#v", i, got[i], r)
		}
	}
}

var (
	red = color.RGBA{0xe1, 0x3b, 0x3b, 0xff}
	blue = color.RGBA{0x3b, 0x9e, 0xe1, 0xff}
	purple = color.RGBA{0x9b, 0x3b, 0xe1, 0xff}
)

var rects = []shape.Interface{
	shape.Rect(0, 0, 1, 1, red),
	shape.Rect(1, 0, 7, 1, blue),
	shape.Rect(8, 0, 1, 1, red),
	shape.Rect(0, 1, 1, 7, blue),
	shape.Rect(1, 1, 3, 1, red),
	shape.Rect(4, 1, 1, 7, purple),
	shape.Rect(5, 1, 3, 1, red),
	shape.Rect(8, 1, 1, 7, blue),
	shape.Rect(2, 2, 2, 1, red),
	shape.Rect(5, 2, 2, 1, red),
	shape.Rect(3, 3, 1, 1, red),
	shape.Rect(5, 3, 1, 1, red),
	shape.Rect(1, 4, 3, 1, purple),
	shape.Rect(5, 4, 3, 1, purple),
	shape.Rect(3, 5, 1, 3, red),
	shape.Rect(5, 5, 1, 3, red),
	shape.Rect(2, 6, 1, 2, red),
	shape.Rect(6, 6, 1, 2, red),
	shape.Rect(1, 7, 1, 1, red),
	shape.Rect(7, 7, 1, 1, red),
	shape.Rect(0, 8, 1, 1, red),
	shape.Rect(1, 8, 7, 1, blue),
	shape.Rect(8, 8, 1, 1, red),
}
