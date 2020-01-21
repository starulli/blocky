package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tinychameleon/blocky/svg"
)

func main() {
	flag.Usage = usage
	debug := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()

	var options []svg.Option
	if *debug {
		options = append(options, svg.DebugMode)
	}

	s, err := svg.New(flag.Arg(0), options...)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Bad input file:", err)
		os.Exit(1)
	}
	fmt.Println(s)
}

func usage() {
	fmt.Printf("Usage: %s [-debug] FILE\n\nFlags:\n", os.Args[0])
	flag.PrintDefaults()
}
