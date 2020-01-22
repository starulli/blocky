// Package shape defines basic shapes for SVG files.
package shape

import (
	"errors"
	"fmt"
	"image/color"
)

// Interface exposes the only useful shape functionality.
type Interface interface {
	fmt.Stringer
	Invisible() bool
	RGBA() color.RGBA
}

// BadHex represents a bad RGBA hex code string.
var BadHex = errors.New("Bad hex code")

// ParseHex converts an RGBA hex colour code with optional alpha into
// an instance of color.RGBA.
func ParseHex(s string) (color.RGBA, error) {
	var r, g, b, a uint8
	n, err := fmt.Sscanf(s, "#%02x%02x%02x%02x", &r, &g, &b, &a)
	if n == 3 {
		err = nil
		a = 0xff
	}
	if err != nil {
		return color.RGBA{}, BadHex
	}
	return color.RGBA{r, g, b, a}, nil
}

func hex(clr color.Color) string {
	c := color.NRGBAModel.Convert(clr).(color.NRGBA)
	return fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
}

func opacity(clr color.Color) string {
	c := color.NRGBAModel.Convert(clr).(color.NRGBA)
	o := float64(c.A) / 255
	v := fmt.Sprintf("%.2f", o)
	if v == "0.00" || v == "1.00" {
		return v[:1]
	}
	return v
}
