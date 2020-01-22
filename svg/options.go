package svg

type Option func(*svg)

func DebugMode(on bool) Option {
	return func(s *svg) {
		s.debug = on
	}
}

func KeepInvisible(on bool) Option {
	return func(s *svg) {
		s.keepInvisible = on
	}
}
