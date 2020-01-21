package shape

import (
	"image/color"
	"testing"
)

func TestPixel(t *testing.T) {
	cases := []struct{
		x, y int
		c color.Color
		want string
	}{
		{0, 0, color.NRGBA{255, 0, 255, 128},
			`<rect width="1" height="1" fill="#ff00ff" fill-opacity="0.50" x="0" y="0"/>`},
		{1, 5, color.NRGBA{},
			`<rect width="1" height="1" fill="#000000" fill-opacity="0" x="1" y="5"/>`},
		{2, 3, color.NRGBA{57, 49, 187, 94},
			`<rect width="1" height="1" fill="#3931bb" fill-opacity="0.37" x="2" y="3"/>`},
		{6, 4, color.NRGBA{0, 0, 0, 255},
			`<rect width="1" height="1" fill="#000000" x="6" y="4"/>`},
	}
	for _, c := range cases {
		got := Pixel(c.x, c.y, c.c).String()
		if got != c.want {
			t.Errorf("Got %v, want %v", got, c.want)
		}
	}
}
