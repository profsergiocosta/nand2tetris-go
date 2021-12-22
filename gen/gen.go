package gen

import (
	"fmt"
	"nand2tetris-go/token"
)

type Generate struct {
	instructions []string
}

func New() *Generate {
	return &Generate{instructions: nil}
}

func (g *Generate) Instructions() []string {
	return g.instructions
}

func (p *Generate) EmmitExpression(tk token.Token) {
	switch tk.Type {
	case token.INT, token.IDENT:
		p.instructions = append(p.instructions, "push")
		p.instructions = append(p.instructions, tk.Lexeme)
	case token.ASTERISK:
		p.instructions = append(p.instructions, "mul")
	case token.PLUS:
		p.instructions = append(p.instructions, "add")

	}
}

func (p *Generate) EmmitAssign(tk token.Token) {
	p.instructions = append(p.instructions, "pop")
	p.instructions = append(p.instructions, tk.Lexeme)
}

func (p *Generate) EmmitPrint() {
	p.instructions = append(p.instructions, "print")
}

func (p *Generate) Disassembly() {

	for i, inst := range p.instructions {
		switch inst {
		case "push", "pop":
			fmt.Printf("(%d)\t%s\t", i, inst)
		case "add", "mul":
			fmt.Printf("(%d)\t%s\t\n", i, inst)
		default:
			fmt.Println(inst)
		}
	}
}
