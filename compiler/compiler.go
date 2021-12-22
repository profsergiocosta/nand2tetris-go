package compiler

import (
	"io/ioutil"
	"nand2tetris-go/gen"
	"nand2tetris-go/parser"
	"nand2tetris-go/translator"
	"nand2tetris-go/vm"
	"path"
	"strings"
)

func FilenameWithoutExtension(fn string) string {
	return strings.TrimSuffix(fn, path.Ext(fn))
}

func Compile(inputPath string) {

	ext := path.Ext(inputPath)
	if ext == ".jack" {
		outputPath := FilenameWithoutExtension(inputPath) + ".vm"

		input, err := ioutil.ReadFile(inputPath)
		if err != nil {
			panic("erro")
		}
		g := gen.New()

		p := parser.New(string(input), g)
		p.Parse()

		vm := vm.New(g.Instructions(), vm.NewEnvironment())
		vm.Save(outputPath)

	} else {
		outputPath := FilenameWithoutExtension(inputPath) + ".asm"
		t := translator.New(inputPath, outputPath)
		t.Translate()
	}
}
