package gen

import (
	"fmt"
	"nand2tetris-go/compiler/token"
)

func GenExpression(tk token.Token) {
	switch tk.Type {
	case token.INT, token.IDENT:
		fmt.Println("push " + tk.Lexeme)
	case token.ASTERISK:
		fmt.Println("mul")
	case token.PLUS:
		fmt.Println("add")

	}
}

func GenAssign(tk token.Token) {
	fmt.Println("pop " + tk.Lexeme)
}
