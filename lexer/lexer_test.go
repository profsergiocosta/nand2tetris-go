package lexer

import (
	"fmt"

	"nand2tetris-go/token"
	"testing"
)

func TestNextToken(t *testing.T) {
	input := "let custo = 80 + 65 * 42;"

	tests := []struct {
		expectedType   token.TokenType
		expectedLexeme string
	}{
		{token.LET, "let"},
		{token.IDENT, "custo"},
		{token.EQ, "="},
		{token.INT, "80"},
		{token.PLUS, "+"},
		{token.INT, "65"},
		{token.ASTERISK, "*"},
		{token.INT, "42"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()
		fmt.Println(tok)

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Lexeme != tt.expectedLexeme {
			t.Fatalf("tests[%d] - lexeme wrong. expected=%q, got=%q",
				i, tt.expectedLexeme, tok.Lexeme)
		}
	}
}
