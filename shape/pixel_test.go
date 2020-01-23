package shape

import (
	"image/color"
	"testing"
)

func TestPixel(t *testing.T) {
	cases := []struct {
		x, y int
		c    color.RGBA
		want string
	}{
		{0, 0, color.RGBA{255, 0, 255, 255},
			`<rect width="1" height="1" fill="#ff00ff" x="0" y="0"/>`},
		{1, 5, color.RGBA{},
			`<rect width="1" height="1" fill="#000000" fill-opacity="0" x="1" y="5"/>`},
		{2, 3, color.RGBA{57, 49, 187, 205},
			`<rect width="1" height="1" fill="#473de9" fill-opacity="0.80" x="2" y="3"/>`},
		{6, 4, color.RGBA{0, 0, 0, 255},
			`<rect width="1" height="1" fill="#000000" x="6" y="4"/>`},
	}
	for _, c := range cases {
		got := Pixel(c.x, c.y, c.c).String()
		if got != c.want {
			t.Errorf("Got %v, want %v", got, c.want)
		}
	}
}

func TestPixelColor(t *testing.T) {
	want := color.RGBA{1, 2, 3, 255}
	p := Pixel(1, 2, color.RGBA{1, 2, 3, 255})
	got := p.RGBA()
	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
