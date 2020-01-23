// Package testdata contains shared data values for all tests.
package testdata

import (
	"encoding/base64"
	"image"
	"io"
	"strings"

	_ "image/png"
)

const ImageData = `
iVBORw0KGgoAAAANSUhEUgAAAAkAAAAJAgMAAACd/+6DAAAAAXNSR0IB2cksfwAAAAlwSFlzAAAL
EwAACxMBAJqcGAAAAAxQTFRF4Ts7O57hmzvhAAAAgavRowAAAAR0Uk5T////AEAqqfQAAAAiSURB
VHicYxANZWBwaHBgKGh2YKjpd2DIWgWhQXyQOFAeAKXYCSUY0KxtAAAAAElFTkSuQmCC
`

func Image() image.Image {
	r := ImageReader()
	m, _, _ := image.Decode(r)
	return m
}

func ImageReader() io.Reader {
	r := strings.NewReader(ImageData)
	return base64.NewDecoder(base64.StdEncoding, r)
}

func BadImageReader() io.Reader {
	r := strings.NewReader(ImageData[9:])
	return base64.NewDecoder(base64.StdEncoding, r)
}
