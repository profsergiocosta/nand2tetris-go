package interpret

import (
	"io/ioutil"
	"nand2tetris-go/gen"
	"nand2tetris-go/parser"
	"nand2tetris-go/vm"
	"path"
)

func Run(filename string) {
	input, err := ioutil.ReadFile(filename)
	if err != nil {
		panic("erro")
	}

	ext := path.Ext(filename)
	if ext == ".jack" {
		gen := gen.New()
		p := parser.New(string(input), gen)
		p.Parse()
		vm := vm.New(gen.Instructions(), vm.NewEnvironment())
		vm.Run()
	} else {
		vm := vm.Open(filename)
		vm.Run()
	}

}
