package parser

import (
	"github.com/matt-tyler/scl/pkg/ast"
	"github.com/matt-tyler/scl/pkg/token"
)

// Parser is used to parse query string expressions
type Parser interface {
	ParseExpr() (ast.Expr, int, error)
	Next() (token.Token, error)
	Peek() (token.Token, error)
}
