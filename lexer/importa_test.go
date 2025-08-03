package lexer

import (
	"testing"

	"github.com/teohen/ttolang/token"
)

func TestImport(t *testing.T) {
	input := `
	importa "./biblioteca.tto";
	`

	tests := []struct {
		typeExpected    token.TokenType
		literalExpected string
	}{
		{token.IMPORT, "importa"},
		{token.STRING, "./biblioteca.tto"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.typeExpected {
			t.Fatalf("tests[%d] - token type wrong. expected=%q, got=%q", i, tt.typeExpected, tok.Literal)
		}
	}
}
