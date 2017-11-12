package ast

import (
	"errors"
	"fmt"
	"time"

	"github.com/matt-tyler/scl/pkg/token"
)

// And is used in an expression tree to represent
// the logical and of two sub expressions.
type and struct {
	left, right Expr
}

// Or is used in an expression tree to represent
// the logical or of two sub expressions.
type or struct {
	left, right Expr
}

func (n and) Eval(t time.Time) bool {
	return n.left.Eval(t) && n.right.Eval(t)
}

func (n or) Eval(t time.Time) bool {
	return n.left.Eval(t) || n.right.Eval(t)
}

// And returns an expression representing the logical
// And of two expressions.
// Returns an error if either expression is nil.
func And(left Expr, right Expr) (Expr, error) {
	if left == nil || right == nil {
		return nil, errors.New("Left or right expression is nil")
	}
	return and{left, right}, nil
}

// Or returns an expression representing the logical
// or of two expressions.
// Returns an error if either expression is nil.
func Or(left Expr, right Expr) (Expr, error) {
	if left == nil || right == nil {
		return nil, errors.New("Left or right expression is nil")
	}
	return or{left, right}, nil
}

// NewBoolExpr returns a new expression representing
// the boolean expression indicated by the tokenType
func NewBoolExpr(op token.TokenType, lhs Expr, rhs Expr) (Expr, error) {
	switch op {
	case token.And:
		e, err := And(lhs, rhs)
		if err != nil {
			return nil, err
		}
		return e, nil
	case token.Or:
		e, err := Or(lhs, rhs)
		if err != nil {
			return nil, err
		}
		return e, nil
	default:
		return nil, fmt.Errorf("Unrecognised operator: %v", op)
	}
}
