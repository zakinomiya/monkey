package lexer

import (
	"monkey/token"
)

type Lexer struct {
  rawText string
  position int
  readPosition int
  char byte
}

func New(input string) *Lexer {
  return &Lexer{
    rawText: input,
    position: 0,
    readPosition: 0,
    char: 0,
  }
}

func (l *Lexer) Size() int {
  return len(l.rawText)
}

func (l *Lexer) isEOF() bool {
  return l.Size() <= l.readPosition
}

func (l *Lexer) skip() bool {
  return l.rawText[l.readPosition] == ' '
}

func (l *Lexer) tokenize() *token.Token {
  s := l.rawText[l.position:l.readPosition] 

  if v, ok := token.TT[s]; ok {
    return token.New(v, s)
  }

  switch s {
	case "ident":
    return token.New(token.Ident, s)
	case "0":
    return token.New(token.Int, s)
  default:
    return token.New(token.Illegal, s)
	}

}

// nextToken reads raw text in the lexer and returns the next token.
func (l *Lexer) nextToken() *token.Token {
  for {
    if l.isEOF() {
      return token.New(token.Eof, "EOF")
    }

    if ok := l.skip(); ok {
      l.position++
      l.readPosition++
      continue
    }

    for !l.isEOF() && !l.skip() {
      l.readPosition++
    }

    tok := l.tokenize()
    l.position = l.readPosition
    return tok
  }
}
