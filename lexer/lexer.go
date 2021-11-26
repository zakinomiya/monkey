package lexer

import "monkey/token"

type Lexer struct {
  rawText string
  position int
  nextPosition int
  char byte
}

func New(input string) *Lexer {
  return &Lexer{
    rawText: input,
    position: 0,
    nextPosition: 1,
    char: 0,
  }
}

func (l *Lexer) nextToken() *token.Token {
  for {
    if l.rawText[l.position: l.position+1] == " " {
      continue
    }

    tok := token.New(l.rawText[l.position:l.position+1])
    l.position++
    return tok
  }
}
