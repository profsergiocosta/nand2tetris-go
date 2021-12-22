package lexer

import (
	"nand2tetris-go/token"
)

type Lexer struct {
	input   string
	current int
	start   int
}

func New(input string) *Lexer {
	l := &Lexer{input: input, current: 0, start: 0}
	return l
}

func (l *Lexer) NextToken() token.Token {

	l.skipWhitespace()
	l.start = l.current

	if l.isAtEnd() {
		return l.makeToken(token.EOF)
	}

	ch := l.peek()

	if token.IsDigit(ch) {
		return l.number()
	}

	if token.IsLetter(ch) {
		return l.identifier()
	}

	switch ch {
	case '=':
		l.advance()
		return l.makeToken(token.EQ)
	case '+':
		l.advance()
		return l.makeToken(token.PLUS)
	case '*':
		l.advance()
		return l.makeToken(token.ASTERISK)
	case ';':
		l.advance()
		return l.makeToken(token.SEMICOLON)

	default:
		l.advance()
		return l.makeToken(token.ILLEGAL)
	}

}

func (l *Lexer) isAtEnd() bool {
	return l.peek() == 0
}

func (l *Lexer) makeToken(ttype token.TokenType) token.Token {
	lexeme := l.input[l.start:l.current]
	return token.Token{Lexeme: lexeme, Type: ttype}

}

func (l *Lexer) number() token.Token {
	for token.IsDigit(l.peek()) {
		l.advance()
	}
	return l.makeToken(token.INT)
}

func (l *Lexer) identifier() token.Token {

	for token.IsLetter(l.peek()) || token.IsDigit(l.peek()) || l.peek() == '_' {
		l.advance()
	}
	tok := l.makeToken(token.IDENT)
	tok.Type = token.LookupIdent(tok.Lexeme)

	return tok
}

func (l *Lexer) peek() byte {
	if l.current >= len(l.input) {
		return 0
	} else {
		return l.input[l.current]
	}
}

func (l *Lexer) advance() byte {
	ch := l.peek()
	if ch != 0 {
		l.current++
	}
	return ch
}

func (l *Lexer) skipWhitespace() {
	ch := l.peek()
	for ch == ' ' || ch == '\t' || ch == '\n' || ch == '\r' {
		l.advance()
		ch = l.peek()
	}
}
