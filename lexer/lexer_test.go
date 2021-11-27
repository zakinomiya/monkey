package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
  input := `= + - * / ( ) { } , ; . let func`

  test := []struct {
    expectedToken token.TokenType
    expectedLiteral string
  }{
    {token.Assign, "="},
    {token.Plus, "+"},
    {token.Minus, "-"},
    {token.Multiply, "*"},
    {token.Division, "/"},
    {token.Lparen, "("},
    {token.Rparen, ")"},
    {token.Lbrace, "{"},
    {token.Rbrace, "}"},
    {token.Comma, ","},
    {token.Semicolon, ";"},
    {token.Period, "."},
    {token.Let, "let"},
    {token.Function, "func"},
  }

  lex := New(input)

  for _, tt := range test {
    tok := lex.nextToken()

    if tok.Type != tt.expectedToken {
      t.Errorf("test failed. expected %s, actual: %s\n", tt.expectedToken.String(), tok.Type)
    }

    if tok.Literal != tt.expectedLiteral {
      t.Errorf("test failed. expected %s, actual: %s\n", tt.expectedLiteral, tok.Literal)
    }
  }
}
