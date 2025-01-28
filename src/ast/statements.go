package ast

import "github.com/hiradsajde/hirad-lang/src/lexer"

type BlockStmt struct {
	Body []Stmt
}

func (b BlockStmt) stmt() {}

type VarDeclarationStmt struct {
	Identifier    string
	Constant      bool
	AssignedValue Expr
	ExplicitType  lexer.Token
}

func (n VarDeclarationStmt) stmt() {}

type ExpressionStmt struct {
	Expression Expr
}

func (n ExpressionStmt) stmt() {}

type Parameter struct {
	Name string
	Type Type
}

type FunctionDeclarationStmt struct {
	Parameters []Parameter
	Name       string
	Body       []Stmt
	ReturnType lexer.Token
}

func (n FunctionDeclarationStmt) stmt() {}

type IfStmt struct {
	Condition  Expr
	Consequent Stmt
}

func (n IfStmt) stmt() {}

type CoutStmt struct {
	Identifier Expr
}

func (n CoutStmt) stmt() {}

type CinStmt struct {
	Identifier Expr
}

func (n CinStmt) stmt() {}

type WhileStmt struct {
	Condition  Expr
	Consequent Stmt
}

func (n WhileStmt) stmt() {}

type ForeachStmt struct {
	Value    string
	Index    bool
	Iterable Expr
	Body     []Stmt
}

func (n ForeachStmt) stmt() {}

type ClassDeclarationStmt struct {
	Name string
	Body []Stmt
}

func (n ClassDeclarationStmt) stmt() {}
