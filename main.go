package main

import (
	"flag"
	"nand2tetris-go/compiler"
	"nand2tetris-go/interpret"
	"nand2tetris-go/repl"
	"os"
)

func main() {

	mode := flag.String("mode", "interpreter", "a string")
	input := flag.String("input", "", "a string")

	flag.Parse()

	if *mode == "compiler" {
		compiler.Compile(*input)
	} else if *mode == "it" {
		repl.Start(os.Stdin, os.Stdout)
	} else {
		interpret.Run(*input)
	}

}
