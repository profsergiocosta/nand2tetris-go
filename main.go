package main

import "nand2tetris-go/compiler/parser"

func main() {
	p := parser.New("let x = 10 * y + 80;")
	p.Parse()
}
