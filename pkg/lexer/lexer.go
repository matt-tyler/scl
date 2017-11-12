package lexer

import (
	"bufio"
	"bytes"
	"io"
	"strings"
	"unicode"

	"github.com/matt-tyler/scl/pkg/token"
)

// Lexer holds the state of the lexer
type Lexer struct {
	s   *bufio.Scanner
	tok token.Token
	pos int
}

func splitter(data []byte, atEOF bool) (int, []byte, error) {
	set := " ,()"
	if len(data) > 0 {
		tBytes := bytes.TrimLeftFunc(data, unicode.IsSpace)
		advance := len(data) - len(tBytes)
		if bytes.ContainsAny(tBytes[0:1], set) { //tBytes[0] == ',' {
			return advance + 1, tBytes[0:1], nil
		}

		pos := bytes.IndexFunc(tBytes, func(c rune) bool {
			return strings.ContainsRune(set, c)
		})

		if pos != -1 {
			b := tBytes[0:pos]
			return advance + len(b), b, nil
		}

		if !atEOF {
			return 0, nil, nil
		}

		return advance + len(tBytes), tBytes, nil
	}
	return 0, nil, nil
}

// New returns a lexer
func New(r io.Reader) *Lexer {
	scanner := bufio.NewScanner(r)
	scanner.Split(splitter)
	return &Lexer{s: scanner}
}

// Next advances the position of the lexer
// returns true if there are still tokens
// available to be consumed
func (l *Lexer) Next() bool {
	l.pos++
	return l.s.Scan()
}

// Token returns the current token pointed
// at by the lexer
func (l *Lexer) Token() token.Token {
	word := l.s.Text()
	tok := Tokenize(word)
	return tok
}

func Tokenize(word string) token.Token {
	var tok token.Token
	switch strings.ToUpper(word) {
	case "ON":
		tok.Typ = token.On
	case "BETWEEN":
		tok.Typ = token.Between
	case "AND":
		tok.Typ = token.And
	case "OR":
		tok.Typ = token.Or
	case "(":
		tok.Typ = token.LeftParen
	case ")":
		tok.Typ = token.RightParen
	case ",":
		tok.Typ = token.Comma
	default:
		tok.Typ = token.String
	}

	if tok.Typ == token.String {
		tok.Val = word
	}

	return tok
}
