package svg

import "image/color"

type Option func(*svg)

func DebugMode(s *svg){
	s.debug = true
}

func KeepInvisible(s *svg) {
	s.keepInvisible = true
}

func Exclude(c color.RGBA) Option {
	return func(s *svg) {
		s.exclude = &c
	}
}
