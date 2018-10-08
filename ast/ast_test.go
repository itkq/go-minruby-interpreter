package ast

import (
	"fmt"
	"testing"

	"github.com/itkq/go-minruby-interpreter/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&LetStatement{
				Token: token.Token{Type: token.LET, Literal: "let"}, // This is a virtual token for parsing
				Name: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "myVar"},
					Value: "myVar",
				},
				Value: &Identifier{
					Token: token.Token{Type: token.IDENT, Literal: "anotherVar"},
					Value: "anotherVar",
				},
			},
		},
	}

	fmt.Printf("str = %+v\n", program.String())
	if program.String() != "myVar = anotherVar;" {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
