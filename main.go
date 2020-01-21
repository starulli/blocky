package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tinychameleon/blocky/svg"
)

func main() {
	flag.Parse()

	s, err := svg.New(flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, "Bad input file:", err)
		os.Exit(1)
	}
	fmt.Println(s)
}
