package parser

import (
	"errors"
	"fmt"

	"github.com/matt-tyler/scl/pkg/ast"
	"github.com/matt-tyler/scl/pkg/token"
)

type pcParser struct {
	pos    int
	tokens []token.Token
}

func (p pcParser) Peek() (token.Token, error) {
	if p.pos >= len(p.tokens) {
		return token.Token{}, fmt.Errorf("No more tokens")
	}
	return p.tokens[p.pos], nil
}

func (p *pcParser) Next() (token.Token, error) {
	if p.pos >= len(p.tokens) {
		return token.Token{}, fmt.Errorf("No more tokens")
	}
	defer func() {
		p.pos++
	}()
	return p.tokens[p.pos], nil
}

// NewPCParser returns a parser that works by using
// precedence climbing
func NewPCParser(tokens []token.Token) Parser {
	return &pcParser{0, tokens}
}

func (p *pcParser) ParseExpr() (ast.Expr, int, error) {
	e, err := p.parseExpr(1)
	return e, p.pos, err
}

func (p *pcParser) parseExpr(prec int) (ast.Expr, error) {
	lhs, err := p.parseAtom()
	if err != nil {
		return nil, err
	}

	for p.pos < len(p.tokens) {
		tok, err := p.Next()
		if err != nil {
			return nil, fmt.Errorf("Unexpected end of input")
		}

		if !token.IsOp(tok.Typ) {
			return nil, fmt.Errorf("Unexpected token: %v", tok)
		}

		if err != nil {
			return nil, fmt.Errorf("An unexpected error occurred: %v", err)
		}

		rhs, err := p.parseExpr(int(tok.Typ) + 1)
		if err != nil {
			return nil, err
		}

		lhs, err = ast.NewBoolExpr(tok.Typ, lhs, rhs)
		if err != nil {
			return nil, err
		}
	}
	return lhs, nil
}

func (p *pcParser) parseAtom() (ast.Expr, error) {
	tok, err := p.Next()
	if err != nil {
		return nil, fmt.Errorf("Unexpected end of input")
	}

	switch tok.Typ {
	case token.LeftParen:
		e, err := p.parseExpr(1)
		if err != nil {
			return nil, err
		}
		tok, err := p.Next()
		if err != nil {
			return nil, fmt.Errorf("Unexpected end of input")
		}

		if tok.Typ != token.RightParen {
			return nil, errors.New("Unmatched \"(\" in input")
		}

		return e, nil
	case token.Between:
		var e Between
		if err := ParseBetween(p, &e); err != nil {
			p.pos--
			return nil, err
		}
		return e, nil
	case token.On:
		var e On
		if err := ParseOn(p, &e); err != nil {
			p.pos--
			return nil, err
		}
		return e, nil
	default:
		p.pos--
		return nil, fmt.Errorf("Unknown token: %v in input", tok)
	}
}
