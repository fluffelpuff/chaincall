package main

import (
	"unicode"
	"unicode/utf8"
)

type Lexer struct {
	input        string
	position     int
	readPosition int
	ch           rune

	line   int
	column int
}

func NewLexer(input string) *Lexer {
	l := &Lexer{input: input, line: 1, column: 0}
	l.readRune()
	return l
}

func (l *Lexer) readRune() {
	if l.readPosition >= len(l.input) {
		l.ch = 0
	} else {
		r, w := utf8.DecodeRuneInString(l.input[l.readPosition:])
		l.ch = r
		l.readPosition += w
	}
	if l.ch == '\n' {
		l.line++
		l.column = 0
	} else {
		l.column++
	}
}

func (l *Lexer) NextToken() Token {
	l.skipWhitespace()

	startLine := l.line
	startCol := l.column

	if l.ch == 0 {
		return Token{Type: EOF, Line: startLine, Column: startCol}
	}

	switch l.ch {
	case ';':
		l.readRune()
		return Token{Type: SEMI, Literal: ";", Line: startLine, Column: startCol}
	case '=':
		l.readRune()
		return Token{Type: ASSIGN, Literal: "=", Line: startLine, Column: startCol}
	case '{':
		l.readRune()
		return Token{Type: LBRACE, Literal: "{", Line: startLine, Column: startCol}
	case '}':
		l.readRune()
		return Token{Type: RBRACE, Literal: "}", Line: startLine, Column: startCol}
	case '(':
		l.readRune()
		return Token{Type: LPAREN, Literal: "(", Line: startLine, Column: startCol}
	case ')':
		l.readRune()
		return Token{Type: RPAREN, Literal: ")", Line: startLine, Column: startCol}
	case '"':
		// String literal
		strVal := l.readString()
		return Token{Type: STRING_LIT, Literal: strVal, Line: startLine, Column: startCol}
	default:
		if isLetter(l.ch) {
			ident := l.readIdentifier()
			ttype := lookupIdent(ident)
			return Token{Type: ttype, Literal: ident, Line: startLine, Column: startCol}
		}
		// Unbekannt
		c := l.ch
		l.readRune()
		return Token{Type: ILLEGAL, Literal: string(c), Line: startLine, Column: startCol}
	}
}

func (l *Lexer) readString() string {
	startPos := l.readPosition
	for {
		l.readRune()
		if l.ch == '"' || l.ch == 0 {
			break
		}
	}
	// substring
	str := l.input[startPos : l.position-1]
	// letztes " konsumieren
	l.readRune()
	return str
}

func (l *Lexer) readIdentifier() string {
	startPos := l.position - 1
	for isLetter(l.ch) {
		l.readRune()
	}
	return l.input[startPos : l.position-1]
}

func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\n' || l.ch == '\r' || l.ch == '\t' {
		l.readRune()
	}
}

func lookupIdent(ident string) TokenType {
	switch ident {
	case "chcall":
		return CHCALL
	case "const":
		return CONSTKW
	case "string":
		return STRINGKW
	case "int":
		return INTKW
	}
	return IDENT
}

func isLetter(ch rune) bool {
	return unicode.IsLetter(ch)
}
