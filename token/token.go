package token

type TokenType int

//go:generate stringer -type=TokenType
const (
	Illegal TokenType = iota
	Eof

	// operators
	Plus
	Minus
	Multiply
	Division

	// symbols
	Assign
	Comma
	Period
	Semicolon
	Lparen
	Rparen
	Lbrace
	Rbrace

	// reserved words
	Function
	Let

	// variable identifier
	Ident

	// values
	Int
)

var TT = map[string]TokenType{
	"=":    Assign,
	"+":    Plus,
	"-":    Minus,
	"*":    Multiply,
	"/":    Division,
	",":    Comma,
	".":    Period,
	";":    Semicolon,
	"(":    Lparen,
	")":    Rparen,
	"{":    Lbrace,
	"}":    Rbrace,
	"func": Function,
	"let":  Let,
}

type Token struct {
	Type    TokenType
	Literal string

	Value interface{}
}

func New(tt TokenType, literal string) *Token {
	return &Token{
		Type:    tt,
		Literal: literal,
	}
}
