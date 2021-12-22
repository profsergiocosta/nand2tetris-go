package parser

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"nand2tetris-go/lexer"
	"nand2tetris-go/token"
	"nand2tetris-go/vm"
	"os"
	"path"
	"strings"
)

type Parser struct {
	lexer    *lexer.Lexer
	curToken token.Token

	instructions []string
}

func New(input string) *Parser {

	lexer := lexer.New(input)
	p := &Parser{lexer: lexer, instructions: nil}
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

func (p *Parser) parseExpression() {
	p.parseTerm()
	for p.curToken.Type == token.ASTERISK || p.curToken.Type == token.PLUS {
		op := p.curToken
		p.nextToken()
		p.parseTerm()
		p.EmmitExpression(op)
	}
}

func (p *Parser) parseTerm() {
	switch p.curToken.Type {
	case token.IDENT:
		p.EmmitExpression(p.curToken)
		p.match(token.IDENT)
	case token.INT:
		p.EmmitExpression(p.curToken)
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
		}

	}
}

func (p *Parser) parsePrintStatement() {
	p.match(token.PRINT)
	p.parseExpression()
	p.match(token.SEMICOLON)
	p.EmmitPrint()
}

func (p *Parser) parseLetStatement() {
	p.match(token.LET)
	ident := p.curToken
	p.match(token.IDENT)
	p.match(token.EQ)
	p.parseExpression()
	p.match(token.SEMICOLON)
	p.EmmitAssign(ident)
}

func (p *Parser) EmmitExpression(tk token.Token) {
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

func (p *Parser) EmmitAssign(tk token.Token) {
	p.instructions = append(p.instructions, "pop")
	p.instructions = append(p.instructions, tk.Lexeme)
}

func (p *Parser) EmmitPrint() {
	p.instructions = append(p.instructions, "print")
}

func (p *Parser) Disassembly() {

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

func Interpret(path string) {
	input, err := ioutil.ReadFile(path)
	if err != nil {
		panic("erro")
	}
	p := New(string(input))
	p.ParseStatements()
	vm := vm.New(p.instructions)
	vm.Run()
}

func FilenameWithoutExtension(fn string) string {
	return strings.TrimSuffix(fn, path.Ext(fn))
}

func Compile(path string) {
	input, err := ioutil.ReadFile(path)
	if err != nil {
		panic("erro")
	}
	p := New(string(input))
	p.ParseStatements()

	f, _ := os.Create(FilenameWithoutExtension(path) + ".vm")

	w := bufio.NewWriter(f)

	for _, inst := range p.instructions {
		w.WriteString(fmt.Sprintf("%s\n", inst))
	}
	w.Flush()
	f.Close()
}
