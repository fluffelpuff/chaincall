package main

import (
	"fmt"
	"strings"
)

type Node interface {
	String() string
}

// Program repräsentiert das komplette Skript
type Program struct {
	Statements []Statement
}

func (p *Program) String() string {
	var sb strings.Builder
	for _, s := range p.Statements {
		sb.WriteString(s.String())
		sb.WriteString("\n")
	}
	return sb.String()
}

// Statement-Interface
type Statement interface {
	Node
	statementNode()
}

// VarDeclaration: z.B. `const string hallo = "welt";`
type VarDeclaration struct {
	IsConst bool
	VarType string
	Name    string
	Value   string // In diesem Minimalbeispiel nur string-Literal (oder Ident?), vereinfacht
}

func (v *VarDeclaration) statementNode() {}
func (v *VarDeclaration) String() string {
	constStr := ""
	if v.IsConst {
		constStr = "const "
	}
	return fmt.Sprintf("%s%s %s = %s;", constStr, v.VarType, v.Name, v.Value)
}

// ChCallStatement: `chcall("url") { ... }`
type ChCallStatement struct {
	URL   string
	Block []Statement // In diesem Block können wieder VarDeclarations usw. sein
}

func (c *ChCallStatement) statementNode() {}
func (c *ChCallStatement) String() string {
	var sb strings.Builder
	sb.WriteString(fmt.Sprintf(`chcall("%s") {`, c.URL))
	if len(c.Block) > 0 {
		sb.WriteString("\n")
		for _, st := range c.Block {
			sb.WriteString("  ")
			sb.WriteString(st.String())
			sb.WriteString("\n")
		}
	}
	sb.WriteString("}")
	return sb.String()
}
