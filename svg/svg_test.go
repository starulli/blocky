package svg

import (
	"github.com/tinychameleon/blocky/grid"
	"github.com/tinychameleon/blocky/shape"
	"github.com/tinychameleon/blocky/testdata"
	"image/color"
	"strings"
	"testing"
)

func TestLoadImage(t *testing.T) {
	m, _ := loadImage(testdata.ImageReader())
	width, height := 9, 9

	if m.Bounds().Max.X != width {
		t.Errorf("Got width %v, want %v", m.Bounds().Max.X, width)
	}

	if m.Bounds().Max.Y != height {
		t.Errorf("Got height %v, want %v", m.Bounds().Max.Y, height)
	}

	m, err := loadImage(testdata.BadImageReader())
	if err == nil {
		t.Errorf("Expected error for bad image")
	}
}

func TestSvgStringerOutput(t *testing.T) {
	m, _ := loadImage(testdata.ImageReader())
	g := grid.New(m)
	s := svg{g: g, debug: false, keepInvisible: false, strategy: fakeStrategy}
	got := s.String()
	if got != svgOutput {
		t.Errorf("Got %v, want %v", got, svgOutput)
	}
}

func fakeStrategy(g grid.Grid) []shape.Interface {
	return []shape.Interface{shape.Pixel(1, 2, color.RGBA{1, 2, 3, 255})}
}

var svgOutput = strings.TrimSpace(`
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 9 9">
<rect width="1" height="1" fill="#010203" x="1" y="2"/>
</svg>
`)
