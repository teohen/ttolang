package ast

import (
	"testing"

	"github.com/teohen/ttolang/token"
)

func TestString(t *testing.T) {
	program := &Program{
		Statements: []Statement{
			&CriaStatement{
				Token: token.Token{Type: token.LET, Literal: "cria"},
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

	expectedOutput := "cria myVar = anotherVar;"

	if program.String() != expectedOutput {
		t.Errorf("program.String() wrong. got=%q", program.String())
	}
}
