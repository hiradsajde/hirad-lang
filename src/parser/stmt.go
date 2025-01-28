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
	symbolName := p.expectError(lexer.IDENTIFIER,
		fmt.Sprintf("Following %s expected variable name however instead recieved %s instead\n",
			lexer.TokenKindString(startToken.Kind), lexer.TokenKindString(p.currentTokenKind())))

	var assignmentValue ast.Expr
	if p.currentTokenKind() != lexer.SEMI_COLON {
		p.expect(lexer.ASSIGNMENT)
		assignmentValue = parse_expr(p, assignment)
		fmt.Println(assignmentValue)
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
	}
	p.expect(lexer.SEMI_COLON)

	if isConstant && assignmentValue == nil {
		panic("Cannot define constant variable without providing default value.")
	}

	return ast.VarDeclarationStmt{
		Constant:      isConstant,
		Identifier:    symbolName.Value,
		AssignedValue: assignmentValue,
		ExplicitType:  startToken,
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
		paramName := p.expect(lexer.IDENTIFIER).Value
		paramType := parse_type(p, defalt_bp)

		functionParams = append(functionParams, ast.Parameter{
			Name: paramName,
			Type: paramType,
		})

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
