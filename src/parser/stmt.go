package parser

import (
	"fmt"

	"github.com/hiradsajde/hirad-lang/src/ast"
	"github.com/hiradsajde/hirad-lang/src/lexer"
)

func parse_stmt(p *parser) ast.Stmt {
	stmt_fn, exists := stmt_lu[p.currentTokenKind()]

	if exists {
		return stmt_fn(p)
	}

	return parse_expression_stmt(p)
}

func parse_expression_stmt(p *parser) ast.ExpressionStmt {
	expression := parse_expr(p, defalt_bp)
	if p.currentTokenKind() == lexer.SEMI_COLON {
		p.expect(lexer.SEMI_COLON)
	}
	return ast.ExpressionStmt{
		Expression: expression,
	}
}

func parse_block_stmt(p *parser) ast.Stmt {
	p.expect(lexer.OPEN_CURLY)
	body := []ast.Stmt{}

	for p.hasTokens() && p.currentTokenKind() != lexer.CLOSE_CURLY {
		body = append(body, parse_stmt(p))
	}

	p.expect(lexer.CLOSE_CURLY)
	return ast.BlockStmt{
		Body: body,
	}
}

func parse_var_decl_stmt(p *parser) ast.Stmt {

	startToken := p.advance()
	if p.nextToken().Kind == lexer.OPEN_PAREN {
		p.dadvance()
		return parse_fn_declaration(p)
	}
	isConstant := startToken.Kind == lexer.CONST
	var assignmentValue ast.Expr
	var results []ast.VarMapDeclarationStmt
	p.advance()
	i := 0
	for p.currentTokenKind() != lexer.SEMI_COLON && p.currentTokenKind() != lexer.EOF {
		symbolName := p.previousToken()
		if p.currentTokenKind() == lexer.COMMA {
			p.advance()
		} else if p.currentTokenKind() == lexer.ASSIGNMENT {
			p.advance()
		} else {
			p.expect(lexer.SEMI_COLON)
		}
		assignmentValue = parse_expr(p, assignment)

		var (
			assignmentType   bool = true
			assignmentString string
		)
		switch assignmentValue.(type) {
		case ast.StringExpr:
			assignmentType = startToken.Kind == lexer.TSTRING
			assignmentString = "string"
		case ast.NumberExpr:
			assignmentType = startToken.Kind == lexer.INT || startToken.Kind == lexer.FLOAT
			assignmentString = "number"
		}
		if !assignmentType {
			panic(fmt.Sprint("variable declaration type is ", startToken.Value, " but ", assignmentString, " assigned"))
		}
		if isConstant && assignmentValue == nil {
			panic("Cannot define constant variable without providing default value.")
		}
		if i%2 == 0 {
			results = append(results, ast.VarMapDeclarationStmt{Identifier: symbolName.Value, AssignedValue: assignmentValue})
		}
		i += 1
	}
	p.expect(lexer.SEMI_COLON)

	return ast.VarDeclarationStmt{
		Constant:     isConstant,
		ExplicitType: startToken,
		Declartion:   results,
	}
}

func parse_fn_params_and_body(p *parser) ([]ast.Parameter, lexer.Token, []ast.Stmt) {
	functionParams := make([]ast.Parameter, 0)
	var returnType lexer.Token
	p.dadvance()
	returnType = p.previousToken()
	p.advance()
	p.expect(lexer.OPEN_PAREN)
	for p.hasTokens() && p.currentTokenKind() != lexer.CLOSE_PAREN {
		if !p.currentToken().IsOneOfMany(lexer.CLOSE_PAREN, lexer.EOF) {
			p.expect(lexer.COMMA)
		}
	}

	p.expect(lexer.CLOSE_PAREN)

	functionBody := ast.ExpectStmt[ast.BlockStmt](parse_block_stmt(p)).Body

	return functionParams, returnType, functionBody
}

func parse_fn_declaration(p *parser) ast.Stmt {
	p.advance()
	functionName := p.expect(lexer.IDENTIFIER).Value
	functionParams, returnType, functionBody := parse_fn_params_and_body(p)

	return ast.FunctionDeclarationStmt{
		Parameters: functionParams,
		ReturnType: returnType,
		Body:       functionBody,
		Name:       functionName,
	}
}

func parse_using_stmt(p *parser) ast.Stmt {
	p.advance()
	if p.currentTokenKind() != lexer.NAMESPACE {
		panic("this compiler only support namespaces `using namespace something`")
	}
	p.advance()
	namespace := parse_expr(p, defalt_bp)
	p.expect(lexer.SEMI_COLON)
	return ast.Namespace{
		Name: namespace,
	}
}

func parse_while_stmt(p *parser) ast.Stmt {
	p.advance()
	condition := parse_expr(p, assignment)
	consequent := parse_block_stmt(p)

	if p.currentTokenKind() == lexer.ELSE {
		p.advance()
	}

	return ast.WhileStmt{
		Condition:  condition,
		Consequent: consequent,
	}
}

func parse_if_stmt(p *parser) ast.Stmt {
	p.advance()
	condition := parse_expr(p, assignment)
	consequent := parse_block_stmt(p)

	if p.currentTokenKind() == lexer.ELSE {
		p.advance()
	}

	return ast.IfStmt{
		Condition:  condition,
		Consequent: consequent,
	}
}

func parse_in_stmt(p *parser) ast.Stmt {
	p.advance()
	p.advance()
	condition := parse_expr(p, assignment)
	p.expect(lexer.SEMI_COLON)
	return ast.CinStmt{
		Identifier: condition,
	}
}

func parse_out_stmt(p *parser) ast.Stmt {
	p.advance()
	p.advance()
	condition := parse_expr(p, assignment)
	p.expect(lexer.SEMI_COLON)
	return ast.CoutStmt{
		Identifier: condition,
	}
}
