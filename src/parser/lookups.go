package parser

import (
	"github.com/hiradsajde/hirad-lang/src/ast"
	"github.com/hiradsajde/hirad-lang/src/lexer"
)

type binding_power int

const (
	defalt_bp binding_power = iota
	comma
	assignment
	logical
	shifting
	relational
	additive
	multiplicative
	unary
	call
	member
	primary
)

type stmt_handler func(p *parser) ast.Stmt
type nud_handler func(p *parser) ast.Expr
type led_handler func(p *parser, left ast.Expr, bp binding_power) ast.Expr

type stmt_lookup map[lexer.TokenKind]stmt_handler
type nud_lookup map[lexer.TokenKind]nud_handler
type led_lookup map[lexer.TokenKind]led_handler
type bp_lookup map[lexer.TokenKind]binding_power

var bp_lu = bp_lookup{}
var nud_lu = nud_lookup{}
var led_lu = led_lookup{}
var stmt_lu = stmt_lookup{}

func led(kind lexer.TokenKind, bp binding_power, led_fn led_handler) {
	bp_lu[kind] = bp
	led_lu[kind] = led_fn
}

func nud(kind lexer.TokenKind, bp binding_power, nud_fn nud_handler) {
	bp_lu[kind] = primary
	nud_lu[kind] = nud_fn
}

func stmt(kind lexer.TokenKind, stmt_fn stmt_handler) {
	bp_lu[kind] = defalt_bp
	stmt_lu[kind] = stmt_fn
}

func createTokenLookups() {

	nud(lexer.INCLUDE, primary, parse_import_expr)

	// Assignment
	led(lexer.ASSIGNMENT, assignment, parse_assignment_expr)
	led(lexer.PLUS_EQUALS, assignment, parse_assignment_expr)
	led(lexer.MINUS_EQUALS, assignment, parse_assignment_expr)
	// Logical
	led(lexer.AND, logical, parse_binary_expr)
	led(lexer.OR, logical, parse_binary_expr)

	// Relational
	led(lexer.LESS, relational, parse_binary_expr)
	led(lexer.LESS_EQUALS, relational, parse_binary_expr)
	led(lexer.GREATER, relational, parse_binary_expr)
	led(lexer.GREATER_EQUALS, relational, parse_binary_expr)
	led(lexer.EQUALS, relational, parse_binary_expr)
	led(lexer.NOT_EQUALS, relational, parse_binary_expr)
	led(lexer.LEFT_SHIFT, relational, parse_binary_expr)
	led(lexer.RIGHT_SHIFT, relational, parse_binary_expr)

	// Additive & Multiplicitave
	led(lexer.PLUS, additive, parse_binary_expr)
	led(lexer.DASH, additive, parse_binary_expr)
	led(lexer.SLASH, multiplicative, parse_binary_expr)
	led(lexer.STAR, multiplicative, parse_binary_expr)
	led(lexer.PERCENT, multiplicative, parse_binary_expr)

	// Literals & Symbols
	nud(lexer.NUMBER, primary, parse_primary_expr)
	nud(lexer.STRING, primary, parse_primary_expr)
	nud(lexer.IDENTIFIER, primary, parse_primary_expr)

	// Unary/Prefix
	nud(lexer.DASH, unary, parse_prefix_expr)
	nud(lexer.NOT, unary, parse_prefix_expr)
	nud(lexer.SHARP, unary, parse_prefix_expr)
	nud(lexer.COUT, unary, parse_prefix_expr)
	nud(lexer.OPEN_BRACKET, primary, parse_array_literal_expr)

	// Member / Computed // Call
	led(lexer.OPEN_BRACKET, member, parse_member_expr)
	led(lexer.OPEN_PAREN, call, parse_call_expr)

	// Grouping Expr
	nud(lexer.OPEN_PAREN, defalt_bp, parse_grouping_expr)

	nud(lexer.INT, defalt_bp, parse_fn_expr)
	nud(lexer.FLOAT, defalt_bp, parse_fn_expr)
	nud(lexer.TSTRING, defalt_bp, parse_fn_expr)
	nud(lexer.BOOL, defalt_bp, parse_fn_expr)
	nud(lexer.RETURN, defalt_bp, parse_return_expr)

	stmt(lexer.USING, parse_using_stmt)

	stmt(lexer.OPEN_CURLY, parse_block_stmt)

	stmt(lexer.INT, parse_fn_declaration)
	stmt(lexer.FLOAT, parse_fn_declaration)
	stmt(lexer.TSTRING, parse_fn_declaration)

	stmt(lexer.INT, parse_var_decl_stmt)
	stmt(lexer.FLOAT, parse_var_decl_stmt)
	stmt(lexer.TSTRING, parse_var_decl_stmt)
	stmt(lexer.CONST, parse_var_decl_stmt)

	stmt(lexer.IF, parse_if_stmt)
	stmt(lexer.CIN, parse_in_stmt)
	stmt(lexer.COUT, parse_out_stmt)
	stmt(lexer.WHILE, parse_while_stmt)
}
