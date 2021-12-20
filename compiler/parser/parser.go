package parser

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"nand2tetris-go/compiler/lexer"
	"nand2tetris-go/compiler/token"
	"nand2tetris-go/vm"
	"os"
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
		p.parseLetStatement()
	}
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

//https://replit.com/@sergio_costa/HonoredSaneTrialsoftware#main.go

func Compile(path string) {
	input, err := ioutil.ReadFile(path)
	if err != nil {
		panic("erro")
	}
	p := New(string(input))
	p.ParseStatements()
	vm := vm.New(p.instructions)

	f, _ := os.Create(path + ".vm")

	w := bufio.NewWriter(f)

	for _, inst := range p.instructions {
		w.WriteString(fmt.Sprintf("%s\n", inst))
	}
	w.Flush()
	f.Close()

	vm.Run()
}
