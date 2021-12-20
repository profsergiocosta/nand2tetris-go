package main

import (
	"nand2tetris-go/compiler/parser"
	"os"
)

func main() {

	fileName := os.Args[1]
	parser.Interpret(fileName)
}
