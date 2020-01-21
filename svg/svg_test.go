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
<rect width="1" height="1" fill="#e13b3b" x="0" y="0"/>
<rect width="1" height="1" fill="#3b9ee1" x="1" y="0"/>
<rect width="1" height="1" fill="#3b9ee1" x="2" y="0"/>
<rect width="1" height="1" fill="#3b9ee1" x="3" y="0"/>
<rect width="1" height="1" fill="#3b9ee1" x="4" y="0"/>
<rect width="1" height="1" fill="#3b9ee1" x="5" y="0"/>
<rect width="1" height="1" fill="#3b9ee1" x="6" y="0"/>
<rect width="1" height="1" fill="#3b9ee1" x="7" y="0"/>
<rect width="1" height="1" fill="#e13b3b" x="8" y="0"/>
<rect width="1" height="1" fill="#3b9ee1" x="0" y="1"/>
<rect width="1" height="1" fill="#e13b3b" x="1" y="1"/>
<rect width="1" height="1" fill="#e13b3b" x="2" y="1"/>
<rect width="1" height="1" fill="#e13b3b" x="3" y="1"/>
<rect width="1" height="1" fill="#9b3be1" x="4" y="1"/>
<rect width="1" height="1" fill="#e13b3b" x="5" y="1"/>
<rect width="1" height="1" fill="#e13b3b" x="6" y="1"/>
<rect width="1" height="1" fill="#e13b3b" x="7" y="1"/>
<rect width="1" height="1" fill="#3b9ee1" x="8" y="1"/>
<rect width="1" height="1" fill="#3b9ee1" x="0" y="2"/>
<rect width="1" height="1" fill="#000000" fill-opacity="0" x="1" y="2"/>
<rect width="1" height="1" fill="#e13b3b" x="2" y="2"/>
<rect width="1" height="1" fill="#e13b3b" x="3" y="2"/>
<rect width="1" height="1" fill="#9b3be1" x="4" y="2"/>
<rect width="1" height="1" fill="#e13b3b" x="5" y="2"/>
<rect width="1" height="1" fill="#e13b3b" x="6" y="2"/>
<rect width="1" height="1" fill="#000000" fill-opacity="0" x="7" y="2"/>
<rect width="1" height="1" fill="#3b9ee1" x="8" y="2"/>
<rect width="1" height="1" fill="#3b9ee1" x="0" y="3"/>
<rect width="1" height="1" fill="#000000" fill-opacity="0" x="1" y="3"/>
<rect width="1" height="1" fill="#000000" fill-opacity="0" x="2" y="3"/>
<rect width="1" height="1" fill="#e13b3b" x="3" y="3"/>
<rect width="1" height="1" fill="#9b3be1" x="4" y="3"/>
<rect width="1" height="1" fill="#e13b3b" x="5" y="3"/>
<rect width="1" height="1" fill="#000000" fill-opacity="0" x="6" y="3"/>
<rect width="1" height="1" fill="#000000" fill-opacity="0" x="7" y="3"/>
<rect width="1" height="1" fill="#3b9ee1" x="8" y="3"/>
<rect width="1" height="1" fill="#3b9ee1" x="0" y="4"/>
<rect width="1" height="1" fill="#9b3be1" x="1" y="4"/>
<rect width="1" height="1" fill="#9b3be1" x="2" y="4"/>
<rect width="1" height="1" fill="#9b3be1" x="3" y="4"/>
<rect width="1" height="1" fill="#9b3be1" x="4" y="4"/>
<rect width="1" height="1" fill="#9b3be1" x="5" y="4"/>
<rect width="1" height="1" fill="#9b3be1" x="6" y="4"/>
<rect width="1" height="1" fill="#9b3be1" x="7" y="4"/>
<rect width="1" height="1" fill="#3b9ee1" x="8" y="4"/>
<rect width="1" height="1" fill="#3b9ee1" x="0" y="5"/>
<rect width="1" height="1" fill="#000000" fill-opacity="0" x="1" y="5"/>
<rect width="1" height="1" fill="#000000" fill-opacity="0" x="2" y="5"/>
<rect width="1" height="1" fill="#e13b3b" x="3" y="5"/>
<rect width="1" height="1" fill="#9b3be1" x="4" y="5"/>
<rect width="1" height="1" fill="#e13b3b" x="5" y="5"/>
<rect width="1" height="1" fill="#000000" fill-opacity="0" x="6" y="5"/>
<rect width="1" height="1" fill="#000000" fill-opacity="0" x="7" y="5"/>
<rect width="1" height="1" fill="#3b9ee1" x="8" y="5"/>
<rect width="1" height="1" fill="#3b9ee1" x="0" y="6"/>
<rect width="1" height="1" fill="#000000" fill-opacity="0" x="1" y="6"/>
<rect width="1" height="1" fill="#e13b3b" x="2" y="6"/>
<rect width="1" height="1" fill="#e13b3b" x="3" y="6"/>
<rect width="1" height="1" fill="#9b3be1" x="4" y="6"/>
<rect width="1" height="1" fill="#e13b3b" x="5" y="6"/>
<rect width="1" height="1" fill="#e13b3b" x="6" y="6"/>
<rect width="1" height="1" fill="#000000" fill-opacity="0" x="7" y="6"/>
<rect width="1" height="1" fill="#3b9ee1" x="8" y="6"/>
<rect width="1" height="1" fill="#3b9ee1" x="0" y="7"/>
<rect width="1" height="1" fill="#e13b3b" x="1" y="7"/>
<rect width="1" height="1" fill="#e13b3b" x="2" y="7"/>
<rect width="1" height="1" fill="#e13b3b" x="3" y="7"/>
<rect width="1" height="1" fill="#9b3be1" x="4" y="7"/>
<rect width="1" height="1" fill="#e13b3b" x="5" y="7"/>
<rect width="1" height="1" fill="#e13b3b" x="6" y="7"/>
<rect width="1" height="1" fill="#e13b3b" x="7" y="7"/>
<rect width="1" height="1" fill="#3b9ee1" x="8" y="7"/>
<rect width="1" height="1" fill="#e13b3b" x="0" y="8"/>
<rect width="1" height="1" fill="#3b9ee1" x="1" y="8"/>
<rect width="1" height="1" fill="#3b9ee1" x="2" y="8"/>
<rect width="1" height="1" fill="#3b9ee1" x="3" y="8"/>
<rect width="1" height="1" fill="#3b9ee1" x="4" y="8"/>
<rect width="1" height="1" fill="#3b9ee1" x="5" y="8"/>
<rect width="1" height="1" fill="#3b9ee1" x="6" y="8"/>
<rect width="1" height="1" fill="#3b9ee1" x="7" y="8"/>
<rect width="1" height="1" fill="#e13b3b" x="8" y="8"/>
</svg>
`)
