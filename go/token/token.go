package token

import "fmt"

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// Identifiers + literals
	IDENTIFIER = "IDENTIFIER" // add, foobar, x, y, ...
	INT        = "INT"        // 1343456

	// Operators
	ASSIGN    = "="
	PLUS      = "+"
	MINUS     = "-"
	BANG      = "!"
	ASTERISK  = "*"
	SLASH     = "/"
	LT        = "<"
	GT        = ">"
	EQUAL     = "=="
	NOT_EQUAL = "!="

	// Delimiters
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// Keywords
	IF   = "IF"
	ELSE = "ELSE"

	TRUE  = "TRUE"
	FALSE = "FALSE"

	FUNCTION = "FUNCTION"
	LET      = "LET"
	RETURN   = "RETURN"
)

type TokenType string

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"if":   IF,
	"else": ELSE,

	"true":  TRUE,
	"false": FALSE,

	"fn":     FUNCTION,
	"let":    LET,
	"return": RETURN,
}

func LookupIdentifier(identifier string) TokenType {
	if token, ok := keywords[identifier]; ok {
		return token
	}

	return IDENTIFIER
}

func Aaa() {
	fmt.Println("Hello from token!")
}
