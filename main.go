package main

import (
	"fmt"
	"io/ioutil"
	"nand2tetris-go/compiler/parser"
	"os"
)

func main() {

	path := os.Args[1]

	input, err := ioutil.ReadFile(path)
	if err != nil {
		panic("erro")
	}
	fmt.Println(string(input))
	p := parser.New(string(input))
	p.Compile()
	p.Disassembly()
}
