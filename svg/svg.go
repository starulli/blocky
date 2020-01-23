// Package svg defines high-level operations for converting image files
// into SVG files.
package svg

import (
	"fmt"
	"github.com/tinychameleon/blocky/grid"
	"github.com/tinychameleon/blocky/strategy"
	"image"
	"image/color"
	"io"
	"os"
	"strings"

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

	s := svg{g: grid.New(m)}
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
	g             grid.Grid
	debug         bool
	keepInvisible bool
	exclude       *color.RGBA
	strategy      strategy.Func
}

func (s svg) String() string {
	var b strings.Builder

	width, height := s.g.Bounds()
	fmt.Fprintf(&b, `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 %d %d">`,
		width, height)

	if s.debug {
		fmt.Fprintln(&b,
			"<style>rect { stroke-width: 0.01; stroke: #000 }</style>")
	}

	fmt.Fprintln(&b)
	for _, sh := range s.strategy(s.g) {
		if !s.keepInvisible && sh.Invisible() || s.exclude != nil && *s.exclude == sh.RGBA() {
			continue
		}
		fmt.Fprintln(&b, sh)
	}

	fmt.Fprintf(&b, "</svg>")
	return b.String()
}
