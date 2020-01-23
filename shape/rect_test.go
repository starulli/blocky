package shape

import (
	"image/color"
	"testing"
)

func TestRect(t *testing.T) {
	cases := []struct {
		x, y, w, h int
		c color.RGBA
		want string
	}{
		{0, 0, 5, 9, color.RGBA{1, 2, 3, 255},
			`<rect width="5" height="9" fill="#010203" x="0" y="0"/>`},
		{1, 3, 2, 1, color.RGBA{},
			`<rect width="2" height="1" fill="#000000" fill-opacity="0" x="1" y="3"/>`},
		{2, 1, 4, 5, color.RGBA{3, 2, 1, 205},
			`<rect width="4" height="5" fill="#030201" fill-opacity="0.80" x="2" y="1"/>`},
	}
	for _, c := range cases {
		got := Rect(c.x, c.y, c.w, c.h, c.c).String()
		if got != c.want {
			t.Errorf("Got %v, want %v", got, c.want)
		}
	}
}

func TestRectColour(t *testing.T) {
	want := color.RGBA{1, 2, 3, 255}
	r := Rect(1, 2, 3, 4, want)
	got := r.RGBA()
	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
