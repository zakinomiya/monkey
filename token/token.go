package token

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identification
	IDENT = "IDENT"
	INT   = "INT"

	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	MULTIPLY = "*"
	DIVISION = "/"

	COMMA     = ","
	PERIOD    = "."
	SEMICOLON = ";"

	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

func New(str string) *Token {
	var literal string
	var tt TokenType

	switch str {
	case EOF:
		literal = "EOF"
		tt = EOF
	case IDENT:
		literal = "IDENT"
		tt = IDENT
	case INT:
		literal = "INT"
		tt = INT
	case ASSIGN:
		literal = "="
		tt = ASSIGN
	case PLUS:
		literal = "+"
		tt = PLUS
	case MINUS:
		literal = "-"
		tt = MINUS
	case MULTIPLY:
		literal = "*"
		tt = MULTIPLY
	case DIVISION:
		literal = "/"
		tt = DIVISION
	case COMMA:
		literal = ","
		tt = COMMA
	case PERIOD:
		literal = "."
		tt = PERIOD
	case SEMICOLON:
		literal = ";"
		tt = SEMICOLON
	case LPAREN:
		literal = "("
		tt = LPAREN
	case RPAREN:
		literal = ")"
		tt = RPAREN
	case LBRACE:
		literal = "{"
		tt = LBRACE
	case RBRACE:
		literal = "}"
		tt = RBRACE
	case FUNCTION:
		literal = "FUNCTION"
		tt = FUNCTION
	case LET:
		literal = "let"
		tt = LET
	}

  return &Token{
    Type: tt,
    Literal: literal,
  }
}
