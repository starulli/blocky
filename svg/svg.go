// Package svg defines high-level operations for converting image files
// into SVG files.
package svg

import (
	"fmt"
	"image"
	"io"
	"os"

	_ "image/png"
)

// New creates an instance of svg.Interface to convert an image
// into an SVG file.
func New(srcFilePath string) (Interface, error) {
	file, err := os.Open(srcFilePath)
	if err != nil {
		return nil, svgError(err)
	}
	defer file.Close()

	m, err := loadImage(file)
	if err != nil {
		return nil, svgError(err)
	}

	return svg{m: m}, nil
}

func loadImage(r io.Reader) (image.Image, error) {
	m, _, err := image.Decode(r)
	return m, err
}

func svgError(e error) error {
	return fmt.Errorf("svg error: %w", e)
}

// Interface exists to explicitly define what can be done to SVG
// instance objects.
type Interface interface {
	fmt.Stringer
}

type svg struct {
	m image.Image
}

func (s svg) String() string {
	return "<svg></svg>"
}
