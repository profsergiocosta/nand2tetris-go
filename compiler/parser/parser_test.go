package parser

import (
	"nand2tetris-go/compiler/token"
	"testing"
)

func TestParseTerm(t *testing.T) {

	tests := []string{"10", "789", "abc", "x"}

	for _, tt := range tests {
		p := New(tt)
		p.parseTerm()
		tk := p.curToken
		if tk.Type != token.EOF {
			t.Fatalf("não encontado fim de arquivo")
		}
	}

}
