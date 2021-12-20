package parser

import (
	"fmt"
	"nand2tetris-go/compiler/gen"
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

func (p *Parser) match(t token.TokenType) {
	if p.curToken.Type == t {
		p.nextToken()
	} else {
		fmt.Println("erro sintático")
		os.Exit(1)
	}
}

func (p *Parser) Parse() {
	p.parseLetStatement()
}

func (p *Parser) parseExpression() {
	p.parseTerm()
	for p.curToken.Type == token.ASTERISK || p.curToken.Type == token.PLUS {
		op := p.curToken
		p.nextToken()
		p.parseTerm()

		gen.GenExpression(op)
	}
}

func (p *Parser) parseTerm() {
	switch p.curToken.Type {
	case token.IDENT:
		{
			gen.GenExpression(p.curToken)
			p.match(token.IDENT)

		}
	case token.INT:
		{
			gen.GenExpression(p.curToken)
			p.match(token.INT)

		}

	default:
		{
			fmt.Println("erro sintático")
			os.Exit(1)
		}
	}

}

func (p *Parser) parseLetStatement() {
	p.match(token.LET)
	ident := p.curToken
	p.match(token.IDENT)
	p.match(token.EQ)
	p.parseExpression()
	p.match(token.SEMICOLON)
	gen.GenAssign(ident)
}
