package shape

import (
	"image/color"
	"testing"
)

func TestParseHex(t *testing.T) {
	cases := []struct{
		in string
		want color.RGBA
		err error
	}{
		{"#000000", color.RGBA{0, 0, 0, 0xff}, nil},
		{"#000000af", color.RGBA{0, 0, 0, 0xaf}, nil},
		{"#12345678", color.RGBA{0x12, 0x34, 0x56, 0x78}, nil},
		{"#1", color.RGBA{}, BadHex},
		{"111", color.RGBA{}, BadHex},
		{"#44", color.RGBA{}, BadHex},
		{"#aaa", color.RGBA{}, BadHex},
	}
	for _, c := range cases {
		got, err := ParseHex(c.in)
		if got != c.want {
			t.Errorf("Got %v, want %v", got, c.want)
		}
		if err != c.err {
			t.Errorf("Got error %v, want %v", err, c.err)
		}
	}

}

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
