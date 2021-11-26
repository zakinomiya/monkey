package lexer

import (
	"monkey/token"
	"testing"
)

func TestNextToken(t *testing.T) {
  input := `=+-*/(){},;.`

  test := []struct {
    expectedToken token.TokenType
    expectedLiteral string
  }{
    {token.ASSIGN, "="},
    {token.PLUS, "+"},
    {token.MINUS, "-"},
    {token.MULTIPLY, "*"},
    {token.DIVISION, "/"},
    {token.LPAREN, "("},
    {token.RPAREN, ")"},
    {token.LBRACE, "{"},
    {token.RBRACE, "}"},
    {token.COMMA, ","},
    {token.SEMICOLON, ";"},
    {token.PERIOD, "."},
  }

  lex := New(input)

  for _, tt := range test {
    tok := lex.nextToken()

    if tok.Type != tt.expectedToken {
      t.Errorf("test failed. expected %s, actual: %s\n", tt.expectedToken, tok.Type)
    }

    if tok.Literal != tt.expectedLiteral {
      t.Errorf("test failed. expected %s, actual: %s\n", tt.expectedLiteral, tok.Literal)
    }
  }
}
