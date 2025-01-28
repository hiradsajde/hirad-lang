package ast

import (
	"github.com/hiradsajde/hirad-lang/src/lexer"
)

// --------------------
// Literal Expressions
// --------------------

type NumberExpr struct {
	Value interface{}
}

func (n NumberExpr) expr() {}

type StringExpr struct {
	Value string
}

func (n StringExpr) expr() {}

type SymbolExpr struct {
	Value string
}

func (n SymbolExpr) expr() {}

// --------------------
// Complex Expressions
// --------------------

type BinaryExpr struct {
	Left     Expr
	Operator lexer.Token
	Right    Expr
}

func (n BinaryExpr) expr() {}

type AssignmentExpr struct {
	Assigne       Expr
	AssignedValue Expr
}

func (n AssignmentExpr) expr() {}

type PrefixExpr struct {
	Operator lexer.Token
	Right    Expr
}

func (n PrefixExpr) expr() {}

type MemberExpr struct {
	Member   Expr
	Property string
}

func (n MemberExpr) expr() {}

type CallExpr struct {
	Method    Expr
	Arguments []Expr
}

func (n CallExpr) expr() {}

type ComputedExpr struct {
	Member   Expr
	Property Expr
}

func (n ComputedExpr) expr() {}

type RangeExpr struct {
	Lower Expr
	Upper Expr
}

func (n RangeExpr) expr() {}

type FunctionExpr struct {
	Parameters []Parameter
	Body       []Stmt
	ReturnType lexer.Token
}

func (n FunctionExpr) expr() {}

type ReturnExpr struct {
	Output interface{}
}

func (n ReturnExpr) expr() {}

type ArrayLiteral struct {
	Contents []Expr
}

func (n ArrayLiteral) expr() {}

type NewExpr struct {
	Instantiation CallExpr
}

func (n NewExpr) expr() {}

type IncludeExpr struct {
	Name string
}

func (n IncludeExpr) expr() {}

type InputExpr struct {
	Inputs []lexer.Token
}

func (b InputExpr) expr() {}
