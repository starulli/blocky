package svg

import "testing"

func TestSvgStringerOutput(t *testing.T) {
	s := svg{}
	got := s.String()
	want := "<svg></svg>"
	if got != want {
		t.Errorf("Got %v, want %v", got, want)
	}
}
