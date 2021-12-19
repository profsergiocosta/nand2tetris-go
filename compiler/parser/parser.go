package parser

import (
	"nand2tetris-go/compiler/lexer"
	"nand2tetris-go/compiler/token"
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
