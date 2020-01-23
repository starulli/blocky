package grid

import (
	"image/color"
	"github.com/tinychameleon/blocky/testdata"
	"testing"
)

func TestNewGrid(t *testing.T) {
	m := testdata.Image()
	g := New(m)

	w, h := g.Bounds()
	max := m.Bounds().Max

	if w != max.X {
		t.Errorf("Got width %v, want %v", w, max.X)
	}

	if h != max.Y {
		t.Errorf("Got height %v, want %v", h, max.Y)
	}

	for y := 0; y < max.Y; y++ {
		for x := 0; x < max.X; x++ {
			got := g.At(x, y)
			want := color.RGBAModel.Convert(m.At(x, y)).(color.RGBA)
			if got != want {
				t.Errorf("Got colour %v, want %v", got, want)
			}
		}
	}
}
