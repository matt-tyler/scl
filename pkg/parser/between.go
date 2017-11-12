package parser

import (
	"fmt"
	"time"

	"github.com/matt-tyler/scl/pkg/token"
)

// Between is an expression that evaluates to true when
// a the time lies between the from and until times
type Between struct {
	from, until time.Time
}

// Eval evaluates the Between expression
func (e Between) Eval(t time.Time) bool {
	return e.from.Before(t) && e.until.After(t)
}

func parseDate(p Parser, t *time.Time) error {
	tok, err := p.Next()
	if err != nil {
		return err
	}

	if tok.Typ != token.String {
		return fmt.Errorf("Unexpected token %v in input", tok)
	}

	format := time.RFC3339

	tt, err := time.Parse(format, tok.Val)
	if err != nil {
		return err
	}

	*t = tt

	return nil
}

// ParseBetween attempts to parse a Between expression
func ParseBetween(p Parser, b *Between) error {
	var x, y time.Time

	if err := parseDate(p, &x); err != nil {
		return err
	}

	if err := parseDate(p, &y); err != nil {
		return err
	}

	*b = Between{x, y}

	return nil
}
