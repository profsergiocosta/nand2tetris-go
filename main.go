package main

import (
	"flag"
	"nand2tetris-go/compiler"
	"nand2tetris-go/interpret"
)

func main() {

	mode := flag.String("mode", "interpreter", "a string")
	input := flag.String("input", "", "a string")

	flag.Parse()

	fileName := *input

	if *mode == "compiler" {
		compiler.Compile(fileName)
	} else {
		interpret.Run(fileName)
	}

}
