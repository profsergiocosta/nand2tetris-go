// token/token.go
package token

type TokenType string

const (
	LET   = "LET"
	IDENT = "IDENT"
	INT   = "INT"

	// OPERADORES
	EQ       = "="
	PLUS     = "+"
	ASTERISK = "*"
)

type Token struct {
	Type   TokenType
	Lexeme string
}
