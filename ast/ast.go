package ast

import "github.com/teohen/ttolang/token"

type Node interface {
	TokenLiteral() string
}

type Statement interface {
	Node
	statementNode()
}

type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	Statements []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

type Identifier struct {
	Token token.Token
	Value string
}

type CriaStatement struct {
	Token token.Token
	Name  *Identifier
	Value Expression
}

func (cs *CriaStatement) statementNode() {}
func (cs *CriaStatement) TokenLiteral() string {
	return cs.Token.Literal
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

type DevolveStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (ds *DevolveStatement) statementNode() {}

func (ds *DevolveStatement) TokenLiteral() string {
	return ds.Token.Literal
}
