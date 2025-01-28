package lexer

import "fmt"

type TokenKind int

const (
	EOF TokenKind = iota
	NULL
	TRUE
	FALSE
	NUMBER
	STRING
	IDENTIFIER

	// Grouping & Braces
	OPEN_BRACKET
	CLOSE_BRACKET
	OPEN_CURLY
	CLOSE_CURLY
	OPEN_PAREN
	CLOSE_PAREN

	// Equivilance
	ASSIGNMENT
	EQUALS
	NOT_EQUALS
	NOT

	// Shorthand
	PLUS_PLUS
	MINUS_MINUS
	PLUS_EQUALS
	MINUS_EQUALS
	NULLISH_ASSIGNMENT // ??=
	LEFT_SHIFT
	RIGHT_SHIFT
	// Conditional
	LESS
	LESS_EQUALS
	GREATER
	GREATER_EQUALS

	// Logical
	OR
	AND

	// Symbols
	SEMI_COLON
	COMMA
	SHARP

	//Maths
	PLUS
	DASH
	SLASH
	STAR
	PERCENT

	// Reserved Keywords
	INT
	TSTRING
	FLOAT
	CONST
	BOOL
	INCLUDE
	USING
	NAMESPACE
	IF
	ELSE
	WHILE
	FOR
	RETURN
	COUT
	CIN
	// Misc
	NUM_TOKENS
)

var reserved_lu map[string]TokenKind = map[string]TokenKind{
	"return":      RETURN,
	"cout":        COUT,
	"cin":         CIN,
	"true":        TRUE,
	"false":       FALSE,
	"null":        NULL,
	"int":         INT,
	"float":       FLOAT,
	"string":      TSTRING,
	"const":       CONST,
	"include":     INCLUDE,
	"if":          IF,
	"else":        ELSE,
	"while":       WHILE,
	"for":         FOR,
	"bool":        BOOL,
	"left_shift":  LEFT_SHIFT,
	"right_shift": RIGHT_SHIFT,
	"using":       USING,
	"namespace":   NAMESPACE,
	"semi_colon":  SEMI_COLON,
}

type Token struct {
	Kind  TokenKind
	Value string
}

func (tk Token) IsOneOfMany(expectedTokens ...TokenKind) bool {
	for _, expected := range expectedTokens {
		if expected == tk.Kind {
			return true
		}
	}

	return false
}

func (token Token) Debug() {
	if token.Kind == IDENTIFIER || token.Kind == NUMBER || token.Kind == STRING {
		fmt.Printf("%s(%s)\n", TokenKindString(token.Kind), token.Value)
	} else {
		fmt.Printf("%s()\n", TokenKindString(token.Kind))
	}
}

func TokenKindString(kind TokenKind) string {
	switch kind {
	case EOF:
		return "eof"
	case NULL:
		return "null"
	case COUT:
		return "cout"
	case CIN:
		return "cin"
	case LEFT_SHIFT:
		return "left_shift"
	case RIGHT_SHIFT:
		return "right_shift"
	case NUMBER:
		return "number"
	case STRING:
		return "string"
	case INT:
		return "int"
	case FLOAT:
		return "float"
	case RETURN:
		return "return"
	case TRUE:
		return "true"
	case FALSE:
		return "false"
	case IDENTIFIER:
		return "identifier"
	case OPEN_BRACKET:
		return "open_bracket"
	case CLOSE_BRACKET:
		return "close_bracket"
	case OPEN_CURLY:
		return "open_curly"
	case CLOSE_CURLY:
		return "close_curly"
	case OPEN_PAREN:
		return "open_paren"
	case CLOSE_PAREN:
		return "close_paren"
	case ASSIGNMENT:
		return "assignment"
	case EQUALS:
		return "equals"
	case NOT_EQUALS:
		return "not_equals"
	case NOT:
		return "not"
	case LESS:
		return "less"
	case LESS_EQUALS:
		return "less_equals"
	case GREATER:
		return "greater"
	case GREATER_EQUALS:
		return "greater_equals"
	case OR:
		return "or"
	case AND:
		return "and"
	case COMMA:
		return "comma"
	case PLUS_PLUS:
		return "plus_plus"
	case MINUS_MINUS:
		return "minus_minus"
	case PLUS_EQUALS:
		return "plus_equals"
	case MINUS_EQUALS:
		return "minus_equals"
	case NULLISH_ASSIGNMENT:
		return "nullish_assignment"
	case PLUS:
		return "plus"
	case DASH:
		return "dash"
	case SLASH:
		return "slash"
	case STAR:
		return "star"
	case PERCENT:
		return "percent"
	case TSTRING:
		return "string"
	case CONST:
		return "const"
	case INCLUDE:
		return "include"
	case SHARP:
		return "sharp"
	case IF:
		return "if"
	case ELSE:
		return "else"
	case FOR:
		return "for"
	case WHILE:
		return "while"
	case BOOL:
		return "bool"
	case USING:
		return "using"
	case NAMESPACE:
		return "namespace"
	case SEMI_COLON:
		return "semi_colon"
	default:
		return fmt.Sprintf("unknown(%d)", kind)
	}
}

func newUniqueToken(kind TokenKind, value string) Token {
	return Token{
		kind, value,
	}
}
