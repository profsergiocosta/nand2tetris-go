package parser

import (
	"nand2tetris-go/token"
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

func TestParseExpression(t *testing.T) {

	tests := []string{"10", "789", "abc", "x", "45+89", "8+6*5"}

	for _, tt := range tests {
		p := New(tt)
		p.parseExpression()
		tk := p.curToken
		if tk.Type != token.EOF {
			t.Fatalf("não encontado fim de arquivo")
		}
	}

}

func TestParseLetStatement(t *testing.T) {

	tests := []string{"let a = 56 * 8 + 9;"}

	for _, tt := range tests {
		p := New(tt)
		p.parseLetStatement()
		tk := p.curToken
		if tk.Type != token.EOF {
			t.Fatalf("não encontado fim de arquivo")
		}
	}

}
