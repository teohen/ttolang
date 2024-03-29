package token

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

const (
	ILLEGAL   = "DESCONHECIDO"
	EOF       = "FIM_DE_ARQUIVO"
	IDENT     = "IDENTIFICADOR"
	INT       = "INT"
	ASSIGN    = "<-"
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	ASTERISK  = "*"
	SLASH     = "/"
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	EQ        = "="
	NEQ       = "!="
	RBRACKET  = "CHAVES_DIR"
	LBRACKET  = "CHAVES_ESQ"

	FUNCTION = "PROCESSO"
	LET      = "CRIACAO"
	VDD      = "VDD"
	FALSE    = "FALSE"
	IF       = "SE"
	ELSE     = "SENAO"
	RETURN   = "DEVOLVE"

	LT = "<"
	GT = ">"

	STRING = "STRING"
)

var keywords = map[string]TokenType{
	"proc":    FUNCTION,
	"cria":    LET,
	"vdd":     VDD,
	"falso":   FALSE,
	"se":      IF,
	"senao":   ELSE,
	"devolve": RETURN,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}

	return IDENT
}
