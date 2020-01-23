package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tinychameleon/blocky/shape"
	"github.com/tinychameleon/blocky/strategy"
	"github.com/tinychameleon/blocky/svg"
)

type cmdFlags struct {
	debug    bool
	keep     bool
	exclude  string
	strategy string
}

func main() {
	var flags cmdFlags
	flag.BoolVar(&flags.debug, "debug", false, "Enable debug mode")
	flag.BoolVar(&flags.keep, "keepInvisible", false,
		"Output elements that have 0x00 alpha values")
	flag.StringVar(&flags.exclude, "exclude", "",
		"Exclude #RRGGBB[AA] colour from output")
	flag.StringVar(&flags.strategy, "optimize", "pixels",
		"Optimize SVG rectangle output. Values: pixels")
	flag.Usage = usage
	flag.Parse()

	options := fromFlags(flags)

	s, err := svg.New(flag.Arg(0), options...)
	if err != nil {
		exit("Bad input file:", err)
	}
	fmt.Println(s)
}

func fromFlags(flags cmdFlags) []svg.Option {
	var o []svg.Option

	if flags.debug {
		o = append(o, svg.DebugMode)
	}

	if flags.keep {
		o = append(o, svg.KeepInvisible)
	}

	if flags.exclude != "" {
		c, err := shape.ParseHex(flags.exclude)
		if err != nil {
			exit("Bad -exclude value:", err)
		}
		o = append(o, svg.Exclude(c))
	}

	o = append(o, svg.Strategy(strategy.PixelForPixel))
	return o
}

func exit(args ...interface{}) {
	fmt.Fprintln(os.Stderr, args...)
	os.Exit(1)
}

func usage() {
	fmt.Println("Usage: %s [-debug] [-keepInvisible] [-exclude=#RRGGBB[AA]] [-optimize=pixels] FILE",
		os.Args[0])
	flag.PrintDefaults()
}
