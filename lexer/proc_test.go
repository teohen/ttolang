package lexer

import (
	"testing"

	"github.com/teohen/ttolang/token"
)

func TestProcs(t *testing.T) {
	input := `
	cria add <- proc(x,y) {
		x + y
	};

	cria result <- add(five, ten);
	{nome <- "ttolang", code <- 1, op <- proc(x) { x; }}
	{op <- proc(x) { x + 2; }}["op"](2);
	`

	tests := []struct {
		typeExpected    token.TokenType
		literalExpected string
	}{
		{token.LET, "cria"},
		{token.IDENT, "add"},
		{token.ASSIGN, "<-"},
		{token.FUNCTION, "proc"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.LET, "cria"},
		{token.IDENT, "result"},
		{token.ASSIGN, "<-"},
		{token.IDENT, "add"},
		{token.LPAREN, "("},
		{token.IDENT, "five"},
		{token.COMMA, ","},
		{token.IDENT, "ten"},
		{token.RPAREN, ")"},
		{token.SEMICOLON, ";"},
		{token.LBRACE, "{"},
		{token.IDENT, "nome"},
		{token.ASSIGN, "<-"},
		{token.STRING, "ttolang"},
		{token.COMMA, ","},
		{token.IDENT, "code"},
		{token.ASSIGN, "<-"},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.IDENT, "op"},
		{token.ASSIGN, "<-"},
		{token.FUNCTION, "proc"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.LBRACE, "{"},
		{token.IDENT, "op"},
		{token.ASSIGN, "<-"},
		{token.FUNCTION, "proc"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.INT, "2"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.RBRACE, "}"},
		{token.LBRACKET, "["},
		{token.STRING, "op"},
		{token.RBRACKET, "]"},
		{token.LPAREN, "("},
		{token.INT, "2"},
		{token.RPAREN, ")"},
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
