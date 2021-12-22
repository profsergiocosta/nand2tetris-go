package parser

import (
	"fmt"
	"nand2tetris-go/gen"
	"nand2tetris-go/lexer"
	"nand2tetris-go/token"
	"os"
)

type Parser struct {
	lexer    *lexer.Lexer
	curToken token.Token

	generate *gen.Generate
}

func New(input string, args ...*gen.Generate) *Parser {

	lexer := lexer.New(input)
	p := &Parser{lexer: lexer}
	if len(args) == 0 {
		p.generate = gen.New()
	} else {
		p.generate = args[0]
	}
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
		fmt.Println(p.curToken)
		os.Exit(1)
	}
}

func (p *Parser) Parse() {
	p.ParseStatements()
}

func (p *Parser) parseExpression() {
	p.parseTerm()
	for p.curToken.Type == token.ASTERISK || p.curToken.Type == token.PLUS {
		op := p.curToken
		p.nextToken()
		p.parseTerm()
		p.generate.EmmitExpression(op)
	}
}

func (p *Parser) parseTerm() {
	switch p.curToken.Type {
	case token.IDENT:
		p.generate.EmmitExpression(p.curToken)
		p.match(token.IDENT)
	case token.INT:
		p.generate.EmmitExpression(p.curToken)
		p.match(token.INT)
	default:
		{
			fmt.Println("erro sintático")
			os.Exit(1)
		}
	}

}

func (p *Parser) ParseStatements() {
	for p.curToken.Type != token.EOF {
		if p.curToken.Type == token.LET {
			p.parseLetStatement()
		} else if p.curToken.Type == token.PRINT {
			p.parsePrintStatement()
		} else { // temporario
			fmt.Println("comando não esperado")
			os.Exit(1)
		}
	}
}

func (p *Parser) parsePrintStatement() {
	p.match(token.PRINT)
	p.parseExpression()
	p.match(token.SEMICOLON)
	p.generate.EmmitPrint()
}

func (p *Parser) parseLetStatement() {
	p.match(token.LET)
	ident := p.curToken
	p.match(token.IDENT)
	p.match(token.EQ)
	p.parseExpression()
	p.match(token.SEMICOLON)
	p.generate.EmmitAssign(ident)
}
