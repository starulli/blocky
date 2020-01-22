package shape

import (
	"fmt"
	"image/color"
)

func Pixel(x int, y int, c color.Color) Interface {
	return pixel{x, y, c}
}

type pixel struct {
	x, y int
	c color.Color
}

func (p pixel) Invisible() bool {
	_, _, _, a := p.c.RGBA()
	return a == 0
}

func (p pixel) RGBA() color.RGBA {
	return color.RGBAModel.Convert(p.c).(color.RGBA)
}

func (p pixel) String() string {
	o := opacity(p.c)
	attr := ""
	if o != "1" {
		attr = fmt.Sprintf(` fill-opacity="%s"`, o)
	}
	return fmt.Sprintf(`<rect width="1" height="1" fill="%s"%s x="%d" y="%d"/>`,
		hex(p.c), attr, p.x, p.y)
}
