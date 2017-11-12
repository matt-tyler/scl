package ast

import "time"

// AST is an abstract syntax tree representing some schedule.
type AST struct {
	allow     bool
	assignees []string
	expr      Expr
}

// Expr represents some function that can be evaluated
// to indicate when something is scheduled.
type Expr interface {
	Eval(time.Time) bool
}
