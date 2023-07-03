package lexer

import "monkey-interpreter/token"

type Lexer struct {
	input        string
	position     int  // current position in input (points to current char)
	readPosition int  // current reading position in input (after current char)
	ch           byte // current char under examination
}

func (lexer *Lexer) readChar() {
	if lexer.readPosition >= len(lexer.input) {
		lexer.ch = 0
	} else {
		lexer.ch = lexer.input[lexer.readPosition]
	}

	lexer.position = lexer.readPosition
	lexer.readPosition += 1
}

func (lexer *Lexer) NextToken() token.Token {
	var tok token.Token

	lexer.skipWhitespace()

	switch lexer.ch {
	case '=':
		tok = newToken(token.ASSIGN, lexer.ch)
	case '+':
		tok = newToken(token.PLUS, lexer.ch)
	case '-':
		tok = newToken(token.MINUS, lexer.ch)
	case '!':
		tok = newToken(token.BANG, lexer.ch)
	case '/':
		tok = newToken(token.SLASH, lexer.ch)
	case '*':
		tok = newToken(token.ASTERISK, lexer.ch)
	case '<':
		tok = newToken(token.LT, lexer.ch)
	case '>':
		tok = newToken(token.GT, lexer.ch)
	case ';':
		tok = newToken(token.SEMICOLON, lexer.ch)
	case '(':
		tok = newToken(token.LPAREN, lexer.ch)
	case ')':
		tok = newToken(token.RPAREN, lexer.ch)
	case ',':
		tok = newToken(token.COMMA, lexer.ch)
	case '{':
		tok = newToken(token.LBRACE, lexer.ch)
	case '}':
		tok = newToken(token.RBRACE, lexer.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		if isLetter(lexer.ch) {
			tok.Literal = lexer.readIdentifier()
			tok.Type = token.LookupIdentifier(tok.Literal)

			return tok
		}

		if isDiggit(lexer.ch) {
			tok.Type = token.INT
			tok.Literal = lexer.readNumber()

			return tok
		}

		tok = newToken(token.ILLEGAL, lexer.ch)
	}

	lexer.readChar()
	return tok
}

func New(input string) *Lexer {
	lexer := &Lexer{input: input}
	lexer.readChar()

	return lexer
}

func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (lexer *Lexer) readIdentifier() string {
	startPosition := lexer.position

	for isLetter(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[startPosition:lexer.position]
}

func (lexer *Lexer) readNumber() string {
	startPosition := lexer.position

	for isDiggit(lexer.ch) {
		lexer.readChar()
	}

	return lexer.input[startPosition:lexer.position]
}

func (lexer *Lexer) skipWhitespace() {
	for lexer.ch == ' ' || lexer.ch == '\t' || lexer.ch == '\r' || lexer.ch == '\n' {
		lexer.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_' || ch == '?' || ch == '\''
}

func isDiggit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
