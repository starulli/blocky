package shape

import (
	"fmt"
	"image/color"
)

func Pixel(x int, y int, c color.RGBA) Interface {
	return pixel{x, y, c}
}

type pixel struct {
	x, y int
	c    color.RGBA
}

func (p pixel) RGBA() color.RGBA {
	return p.c
}

func (p pixel) String() string {
	return fmt.Sprintf(`<rect width="1" height="1" fill="%s"%s x="%d" y="%d"/>`,
		hex(p.c), opacityAttr(p.c), p.x, p.y)
}
