package lexer

import "github.com/brend-n/monkey-interpreter/token"

type Lexer struct {
	input        string
	position     int // points to current char
	readPosition int // points to reading position (next char)
	ch           byte
}

func New(input string) *Lexer {
	l := &Lexer{
		input: input,
	}
	l.readChar()
	return l
}

func (l *Lexer) readChar() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	l.position = l.readPosition
	l.readPosition++
}

func (l *Lexer) NextToken() token.Token {
	l.consumeWhitespace()
	tok := token.Token{Literal: string(l.ch)}

	switch l.ch {
	case '=':
		tok.Type = token.ASSIGN
	case '+':
		tok.Type = token.PLUS
	case '(':
		tok.Type = token.LPAREN
	case ')':
		tok.Type = token.RPAREN
	case '{':
		tok.Type = token.LBRACE
	case '}':
		tok.Type = token.RBRACE
	case ';':
		tok.Type = token.SEMICOLON
	case ',':
		tok.Type = token.COMMA
	case 0:
		tok.Type = token.EOF
		tok.Literal = ""
	default:
		if isLetter(l.ch) {
			tok.Literal = l.readValue(isLetter)
			tok.Type = token.LookupIdent(tok.Literal)
			return tok
		} else if isDigit(l.ch) {
			tok.Literal = l.readValue(isDigit)
			tok.Type = token.INT
			return tok
		} else {
			tok.Type = token.ILLEGAL
		}
	}

	l.readChar()

	return tok
}

func (l *Lexer) readValue(cb func(byte) bool) string {
	position := l.position
	for cb(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

func (l *Lexer) consumeWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}
