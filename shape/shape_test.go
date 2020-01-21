package shape

import (
	"image/color"
	"testing"
)

func TestHex(t *testing.T) {
	cases := []struct{
		in color.Color
		want string
	}{
		{color.NRGBA{}, "#000000"},
		{color.NRGBA{255, 255, 255, 255}, "#ffffff"},
		{color.NRGBA{45, 89, 194, 13}, "#2d59c2"},
	}
	for _, c := range cases {
		got := hex(c.in)
		if got != c.want {
			t.Errorf("Got %v, want %v", got, c.want)
		}
	}
}

func TestOpacity(t *testing.T) {
	cases := []struct{
		in color.Color
		want string
	}{
		{color.NRGBA{}, "0"},
		{color.NRGBA{255, 255, 255, 255}, "1"},
		{color.NRGBA{45, 89, 194, 13}, "0.05"},
		{color.NRGBA{123, 62, 29, 128}, "0.50"},
	}
	for _, c := range cases {
		got := opacity(c.in)
		if got != c.want {
			t.Errorf("Got %v, want %v", got, c.want)
		}
	}
}
