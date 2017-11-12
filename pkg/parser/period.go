package parser

import (
	"fmt"
	"strings"
	"time"

	"github.com/matt-tyler/scl/pkg/token"
)

// Period is used to represent a
// particular period
type Period int

const (
	sundays Period = iota
	mondays
	tuesdays
	wednesdays
	thursdays
	fridays
	saturdays
	weekdays
	weekends
)

// On is an expression representing a list of periods
// that can be tested against to determine if a particular
// time is part of the schedule
type On struct {
	periods []Period
}

// ParseOn parses an On expression
func ParseOn(p Parser, o *On) error {
	periods := []Period{}
	var period Period
	if err := ParsePeriod(p, &period); err != nil {
		return err
	}

	periods = append(periods, period)

	tok, err := p.Peek()
	for ; err == nil && tok.Typ == token.Comma; tok, err = p.Peek() {
		_, err := p.Next()
		if err != nil {
			return fmt.Errorf("Unexpected error parsing %v", tok)
		}

		if err := ParsePeriod(p, &period); err != nil {
			return err
		}
		periods = append(periods, period)
	}

	o.periods = periods
	return nil
}

// TokenToPeriod converts a token to a period
func TokenToPeriod(tok token.Token, period *Period) error {
	switch strings.ToUpper(tok.Val) {
	case "MONDAYS":
		*period = mondays
	case "TUESDAYS":
		*period = tuesdays
	case "WEDNESDAYS":
		*period = wednesdays
	case "THURSDAYS":
		*period = thursdays
	case "FRIDAYS":
		*period = fridays
	case "SATURDAYS":
		*period = saturdays
	case "SUNDAYS":
		*period = sundays
	case "WEEKDAYS":
		*period = weekdays
	case "WEEKENDS":
		*period = weekends
	default:
		return fmt.Errorf("%v is not valid period literal", tok)
	}

	return nil
}

// ParsePeriod does stuff
func ParsePeriod(p Parser, period *Period) error {
	tok, err := p.Next()
	if err != nil {
		return fmt.Errorf("Expected period literal: %v", err)
	}

	if tok.Typ != token.String {
		return fmt.Errorf("Expected period literal, got %v instead", tok)
	}

	return TokenToPeriod(tok, period)
}

func isWeekday(w time.Weekday) bool {
	return time.Monday <= w && w <= time.Friday
}

// Eval determines whether the time is included
// in the specified periods
func (s On) Eval(t time.Time) bool {
	day := t.Weekday()

	for _, period := range s.periods {
		if int(period) == int(day) {
			return true
		}
		if period == weekdays && isWeekday(day) {
			return true
		}
		if period == weekends && !isWeekday(day) {
			return true
		}
	}
	return false
}

// ToPeriod converts a string representation of
// a period to its enumerated type version
func ToPeriod(s string) (Period, error) {
	switch s {
	case "MONDAYS":
		return mondays, nil
	case "TUESDAYS":
		return tuesdays, nil
	case "WEDNESDAYS":
		return wednesdays, nil
	case "THURSDAYS":
		return thursdays, nil
	case "FRIDAYS":
		return fridays, nil
	case "SATURDAYS":
		return saturdays, nil
	case "SUNDAYS":
		return sundays, nil
	case "WEEKDAYS":
		return weekdays, nil
	case "WEEKENDS":
		return weekends, nil
	default:
		return 0, fmt.Errorf("%v is not valid period literal", s)
	}
}
