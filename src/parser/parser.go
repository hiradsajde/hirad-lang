package parser

import (
	"github.com/hiradsajde/hirad-lang/src/ast"
	"github.com/hiradsajde/hirad-lang/src/lexer"
)

type parser struct {
	tokens []lexer.Token
	pos    int
}

func createParser(tokens []lexer.Token) *parser {
	p := &parser{
		tokens: tokens,
		pos:    0,
	}

	return p
}

func Parse(source string) ast.BlockStmt {
	tokens := lexer.Tokenize(source)
	p := createParser(tokens)
	body := make([]ast.Stmt, 0)

	for p.hasTokens() {
		body = append(body, parse_stmt(p))
	}

	return ast.BlockStmt{
		Body: body,
	}
}
