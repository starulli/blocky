// Package svg defines high-level operations for converting image files
// into SVG files.
package svg

import (
	"fmt"
)

// New creates an instance of svg.Interface to convert an image
// into an SVG file.
func New(srcPath string) (Interface, error) {
	return svg{}, nil
}

// Interface exists to explicitly define what can be done to SVG
// instance objects.
type Interface interface {
	fmt.Stringer
}

type svg struct {
}

func (s svg) String() string {
	return "<svg></svg>"
}
