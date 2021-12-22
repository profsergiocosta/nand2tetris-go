package lexer

import (
	"fmt"
	"nand2tetris-go/compiler/token"
)

type Lexer struct {
	input    string
	position int
	start    int
}

func New(input string) *Lexer {
	l := &Lexer{input: input, position: 0, start: 0}
	return l
}

func (l *Lexer) NextToken() token.Token {

	l.skipWhitespace()

	l.start = l.position

	ch := l.peekChar()

	switch ch {
	case '=':
		return token.Token{Type: token.EQ, Lexeme: string(l.nextChar())}
	case '+':
		return token.Token{Type: token.PLUS, Lexeme: string(l.nextChar())}
	case '*':
		return token.Token{Type: token.ASTERISK, Lexeme: string(l.nextChar())}
	case ';':
		return token.Token{Type: token.SEMICOLON, Lexeme: string(l.nextChar())}

	case 0:
		return token.Token{Type: token.EOF, Lexeme: ""}

	default:
		if token.IsDigit(ch) {
			return l.readInt()
		} else if token.IsLetter(ch) {
			return l.readIdentifier()
		} else {
			fmt.Println(ch)
			return token.Token{Type: token.ILLEGAL, Lexeme: string(l.nextChar())}
		}

	}

}

func (l *Lexer) makeToken(ttype token.TokenType) token.Token {
	lexeme := l.input[l.start:l.position]
	return token.Token{Lexeme: lexeme, Type: ttype}

}

func (l *Lexer) readInt() token.Token {
	//var tok token.Token
	//tok.Type = token.INT

	//position := l.position
	for token.IsDigit(l.peekChar()) {
		l.nextChar()
	}
	//lexeme := l.input[position:l.position]
	//tok.Lexeme = lexeme

	//return tok
	return l.makeToken(token.INT)
}

func (l *Lexer) readIdentifier() token.Token {
	//var tok token.Token

	//position := l.position

	for token.IsLetter(l.peekChar()) || token.IsDigit(l.peekChar()) || l.peekChar() == '_' {
		l.nextChar()
	}

	//lexeme := l.input[position:l.position]
	/*
		if lexeme == "let" {
			tok.Type = token.LET
		} else if lexeme == "print" {
			tok.Type = token.PRINT
		} else {
			tok.Type = token.IDENT
		}
		tok.Lexeme = lexeme
		return tok
	*/
	tok := l.makeToken(token.IDENT)
	tok.Type = token.LookupIdent(tok.Lexeme)

	return tok
}

func (l *Lexer) peekChar() byte {
	if l.position >= len(l.input) {
		return 0
	} else {
		return l.input[l.position]
	}
}

func (l *Lexer) nextChar() byte {
	ch := l.peekChar()
	if ch != 0 {
		l.position++
	}
	return ch
}

func (l *Lexer) skipWhitespace() {
	ch := l.peekChar()
	for ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
		l.nextChar()
		ch = l.peekChar()
	}
}
