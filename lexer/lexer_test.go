package lexer

import (
	"testing"

	"github.com/teohen/ttolang/token"
)

func TestNextToken(t *testing.T) {
	input := `
	cria five <- 5;
	cria ten <- 10;

	cria add <- proc(x,y) {
		x + y
	};

	cria result <- add(five, ten);
	!-/*5;
	5 < 10 > 5;

	se (5 < 10) {
		devolve vdd;
	} senao {
		devolve falso;
	}

	10 = 10;
	10 != 9;
	"foobar"
	"foo bar"
	[1, 2];
	`

	tests := []struct {
		typeExpected    token.TokenType
		literalExpected string
	}{
		{token.LET, "cria"},
		{token.IDENT, "five"},
		{token.ASSIGN, "<-"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.LET, "cria"},
		{token.IDENT, "ten"},
		{token.ASSIGN, "<-"},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
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
		{token.BANG, "!"},
		{token.MINUS, "-"},
		{token.SLASH, "/"},
		{token.ASTERISK, "*"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.GT, ">"},
		{token.INT, "5"},
		{token.SEMICOLON, ";"},
		{token.IF, "se"},
		{token.LPAREN, "("},
		{token.INT, "5"},
		{token.LT, "<"},
		{token.INT, "10"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RETURN, "devolve"},
		{token.VDD, "vdd"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.ELSE, "senao"},
		{token.LBRACE, "{"},
		{token.RETURN, "devolve"},
		{token.FALSE, "falso"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.INT, "10"},
		{token.EQ, "="},
		{token.INT, "10"},
		{token.SEMICOLON, ";"},
		{token.INT, "10"},
		{token.NEQ, "!="},
		{token.INT, "9"},
		{token.SEMICOLON, ";"},
		{token.STRING, "foobar"},
		{token.STRING, "foo bar"},
		{token.LBRACKET, "["},
		{token.INT, "1"},
		{token.COMMA, ","},
		{token.INT, "2"},
		{token.RBRACKET, "]"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.typeExpected {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.typeExpected, tok.Literal)
		}
	}
}
