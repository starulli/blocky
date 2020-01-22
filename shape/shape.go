// Package shape defines basic shapes for SVG files.
package shape

import (
	"fmt"
	"image/color"
)

// Interface exposes the only useful shape functionality.
type Interface interface {
	fmt.Stringer
	Invisible() bool
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
