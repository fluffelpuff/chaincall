package main

import "fmt"

type Parser struct {
	lexer     *Lexer
	curToken  Token
	peekToken Token
	errors    []string
}

func NewParser(l *Lexer) *Parser {
	p := &Parser{lexer: l}
	p.nextToken()
	p.nextToken()
	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.lexer.NextToken()
}

func (p *Parser) ParseProgram() *Program {
	prog := &Program{}
	for p.curToken.Type != EOF && p.curToken.Type != ILLEGAL {
		stmt := p.parseStatement()
		if stmt != nil {
			prog.Statements = append(prog.Statements, stmt)
		}
		p.nextToken()
	}
	return prog
}

func (p *Parser) parseStatement() Statement {
	switch p.curToken.Type {
	case CONSTKW, STRINGKW, INTKW:
		// VarDeclaration
		return p.parseVarDeclaration()
	case CHCALL:
		// chcall("url") { block }
		return p.parseChCallStatement()
	default:
		// Unbekannt oder wir ignorieren
		return nil
	}
}

func (p *Parser) parseVarDeclaration() Statement {
	// optional const
	isConst := false
	if p.curToken.Type == CONSTKW {
		isConst = true
		p.nextToken() // nun sollte z.B. string/int kommen
	}

	varType := p.curToken.Literal // e.g. "string", "int"
	p.nextToken()                 // expect IDENT
	if p.curToken.Type != IDENT {
		p.errorf("Expected identifier, got '%s'", p.curToken.Literal)
		return nil
	}
	varName := p.curToken.Literal

	p.nextToken() // expect '='
	if p.curToken.Type != ASSIGN {
		p.errorf("Expected '=', got '%s'", p.curToken.Literal)
		return nil
	}
	// value -> wir tun so, als wären es nur strings oder ident
	p.nextToken()
	val := p.curToken.Literal

	// optional semicolon
	if p.peekToken.Type == SEMI {
		p.nextToken() // semicolon
	}

	return &VarDeclaration{
		IsConst: isConst,
		VarType: varType,
		Name:    varName,
		Value:   val,
	}
}

func (p *Parser) parseChCallStatement() Statement {
	// Wir sind auf CHCALL
	p.nextToken() // sollte '(' sein
	if p.curToken.Type != LPAREN {
		p.errorf("Expected '(' after chcall, got '%s'", p.curToken.Literal)
		return nil
	}
	p.nextToken() // erwartet STRING_LIT
	if p.curToken.Type != STRING_LIT {
		p.errorf("Expected string literal for URL, got '%s'", p.curToken.Literal)
		return nil
	}
	url := p.curToken.Literal

	p.nextToken() // sollte ')'
	if p.curToken.Type != RPAREN {
		p.errorf("Expected ')' after URL, got '%s'", p.curToken.Literal)
		return nil
	}

	p.nextToken() // sollte '{'
	if p.curToken.Type != LBRACE {
		p.errorf("Expected '{' after chcall(...)")
		return nil
	}

	// parse block statements bis '}'
	blockStmts := []Statement{}
	p.nextToken() // erstes Statement oder '}'?
	for p.curToken.Type != RBRACE && p.curToken.Type != EOF {
		stmt := p.parseStatement()
		if stmt != nil {
			blockStmts = append(blockStmts, stmt)
		}
		p.nextToken()
	}

	// now we consume '}'
	chSt := &ChCallStatement{
		URL:   url,
		Block: blockStmts,
	}
	return chSt
}

func (p *Parser) errorf(format string, args ...interface{}) {
	msg := fmt.Sprintf(format, args...)
	p.errors = append(p.errors, msg)
}
