package svg

import (
	"encoding/base64"
	"io"
	"strings"
	"testing"
)

func TestLoadImage(t *testing.T) {
	m, _ := loadImage(imageReader())
	width, height := 9, 9

	if m.Bounds().Max.X != width {
		t.Errorf("Got width %v, want %v", m.Bounds().Max.X, width)
	}

	if m.Bounds().Max.Y != height {
		t.Errorf("Got height %v, want %v", m.Bounds().Max.Y, height)
	}

	m, err := loadImage(badImageReader())
	if err == nil {
		t.Errorf("Expected error for bad image")
	}
}

func TestSvgStringerOutput(t *testing.T) {
	m, _ := loadImage(imageReader())
	s := svg{m: m}
	got := s.String()
	if got != svgOutput {
		t.Errorf("Got %v, want %v", got, svgOutput)
	}
}

func imageReader() io.Reader {
	return base64.NewDecoder(base64.StdEncoding,
		strings.NewReader(imageData))
}

func badImageReader() io.Reader {
	badData := imageData[9:]
	return base64.NewDecoder(base64.StdEncoding,
		strings.NewReader(badData))
}

const imageData = `
iVBORw0KGgoAAAANSUhEUgAAAAkAAAAJAgMAAACd/+6DAAAAAXNSR0IB2cksfwAAAAlwSFlzAAAL
EwAACxMBAJqcGAAAAAxQTFRF4Ts7O57hmzvhAAAAgavRowAAAAR0Uk5T////AEAqqfQAAAAiSURB
VHicYxANZWBwaHBgKGh2YKjpd2DIWgWhQXyQOFAeAKXYCSUY0KxtAAAAAElFTkSuQmCC
`

var svgOutput = strings.TrimSpace(`
<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 9 9">
</svg>
`)
