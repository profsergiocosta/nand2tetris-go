package parser

import (
	"fmt"
	"nand2tetris-go/compiler/lexer"
	"nand2tetris-go/compiler/token"
	"os"
)

type Parser struct {
	lexer    *lexer.Lexer
	curToken token.Token
}

func New(input string) *Parser {

	lexer := lexer.New(input)
	p := &Parser{lexer: lexer}
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.lexer.NextToken()
}

func (p *Parser) expect(t token.TokenType) {
	if p.curToken.Type == t {
		p.nextToken()
	} else {
		fmt.Println("erro sintático")
		os.Exit(1)
	}
}

func (p *Parser) parseTerm() {
	switch p.curToken.Type {
	case token.IDENT:
		p.expect(token.IDENT)
	case token.INT:
		p.expect(token.INT)
	default:
		{
			fmt.Println("erro sintático")
			os.Exit(1)
		}
	}

}
