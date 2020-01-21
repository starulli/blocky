package svg

type Option func(*svg)

func DebugMode(s *svg) {
	s.debug = true
}
