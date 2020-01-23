package shape

import (
	"fmt"
	"image/color"
)

func Rect(x int, y int, w int, h int, c color.RGBA) Interface {
	return rect{x, y, w, h, c}
}

type rect struct {
	x, y, w, h int
	c color.RGBA
}

func (r rect) RGBA() color.RGBA {
	return r.c
}

func (r rect) String() string {
	return fmt.Sprintf(
		`<rect width="%d" height="%d" fill="%s"%s x="%d" y="%d"/>`,
		r.w, r.h, hex(r.c), opacityAttr(r.c), r.x, r.y)
}
