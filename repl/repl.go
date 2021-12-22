package repl

import (
	"bufio"
	"fmt"
	"io"
	"nand2tetris-go/gen"
	"nand2tetris-go/parser"
	"nand2tetris-go/vm"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := vm.NewEnvironment()

	for {
		fmt.Print(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text() + ";"

		gen := gen.New()
		p := parser.New(line, gen)
		p.Parse()
		vm := vm.New(gen.Instructions(), env)
		vm.Run()
	}
}
