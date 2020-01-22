// Package svg defines high-level operations for converting image files
// into SVG files.
package svg

import (
	"fmt"
	"image"
	"io"
	"os"
	"strings"
	"github.com/tinychameleon/blocky/shape"

	_ "image/png"
)

// New creates an instance of svg.Interface to convert an image
// into an SVG file.
func New(srcFilePath string, options ...Option) (Interface, error) {
	file, err := os.Open(srcFilePath)
	if err != nil {
		return nil, svgError(err)
	}
	defer file.Close()

	m, err := loadImage(file)
	if err != nil {
		return nil, svgError(err)
	}

	s := svg{m: m}
	for _, opt := range options {
		opt(&s)
	}
	return s, nil
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
	debug bool
	keepInvisible bool
}

func (s svg) convert() []shape.Interface {
	var sh []shape.Interface
	for y := 0; y < s.m.Bounds().Max.Y; y++ {
		for x := 0; x < s.m.Bounds().Max.X; x++ {
			sh = append(sh, shape.Pixel(x, y, s.m.At(x, y)))
		}
	}
	return sh
}

func (s svg) String() string {
	var b strings.Builder

	fmt.Fprintf(&b, `<svg xmlns="http://www.w3.org/2000/svg" `)
	fmt.Fprintf(&b, `viewBox="0 0 %d %d">`,
		s.m.Bounds().Max.X, s.m.Bounds().Max.Y)

	if s.debug {
		fmt.Fprintln(&b,
			"<style>rect { stroke-width: 0.01; stroke: #000 }</style>")
	}

	fmt.Fprintln(&b)
	for _, sh := range s.convert() {
		if !s.keepInvisible && sh.Invisible() {
			continue
		}
		fmt.Fprintln(&b, sh)
	}

	fmt.Fprintf(&b, "</svg>")
	return b.String()
}
