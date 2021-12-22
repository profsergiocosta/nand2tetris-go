// token/token.go
package token

type TokenType string

const (
	EOF     = "EOF"
	ILLEGAL = "ILLEGAL"

	LET   = "LET"
	IDENT = "IDENT"
	INT   = "INT"
	PRINT = "PRINT"

	// OPERADORES
	EQ        = "="
	PLUS      = "+"
	ASTERISK  = "*"
	SEMICOLON = ";"
)

var keywords = map[string]TokenType{
	"let":   LET,
	"print": PRINT,
}

type Token struct {
	Type   TokenType
	Lexeme string
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}

func IsLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func IsDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
