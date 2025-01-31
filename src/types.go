package main

type TokenType string

const (
	EOF        TokenType = "EOF"
	ILLEGAL    TokenType = "ILLEGAL"
	IDENT      TokenType = "IDENT"
	SEMI       TokenType = ";"
	ASSIGN     TokenType = "="
	LBRACE     TokenType = "{"
	RBRACE     TokenType = "}"
	LPAREN     TokenType = "("
	RPAREN     TokenType = ")"
	STRING_LIT TokenType = "STRING_LIT"

	// Neues Keyword
	CHCALL TokenType = "CHCALL"

	// Mögliche Typen
	STRINGKW TokenType = "string"
	INTKW    TokenType = "int"
	CONSTKW  TokenType = "const"
)

type Token struct {
	Type    TokenType
	Literal string
	Line    int
	Column  int
}
