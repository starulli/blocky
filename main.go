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
	keep := flag.Bool("keepInvisible", false,
		"Output elements that have 0x00 alpha values")
	flag.Parse()

	s, err := svg.New(flag.Arg(0), svg.DebugMode(*debug),
		svg.KeepInvisible(*keep))
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
