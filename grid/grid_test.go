package grid

import (
	"github.com/tinychameleon/blocky/testdata"
	"image/color"
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
			if *got != want {
				t.Errorf("Got colour %v, want %v", *got, want)
			}
		}
	}
}

func TestGridErase(t *testing.T) {
	x, y := 0, 0
	g := New(testdata.Image())
	if g.At(x, y) == nil {
		t.Fatalf("g.At(%d, %d) is already empty", x, y)
	}
	g.EraseAt(x, y)
	got := g.At(x, y)
	if got != nil {
		t.Errorf("Got %v, want nil", got)
	}
}

func TestGridEraseMatching(t *testing.T) {
	g := New(testdata.Image())
	c := *g.At(0, 0)
	g.EraseMatching(func(x color.RGBA) bool {
		return x == c
	})

	width, height := g.Bounds()
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			got := g.At(x, y)
			if got != nil && *got == c {
				t.Errorf("Found erased colour at (%d, %d)", x, y)
			}
		}
	}
}
