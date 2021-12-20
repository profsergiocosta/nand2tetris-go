package main

import (
	"nand2tetris-go/compiler/vm"
	"os"
)

func main() {

	fileName := os.Args[1]
	//parser.Interpret(fileName)
	//parser.Compile(fileName)
	vm.Interpret(fileName)
}
