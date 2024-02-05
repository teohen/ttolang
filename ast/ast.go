package ast

import (
	"bytes"
	"strings"

	"github.com/teohen/ttolang/token"
)

type Node interface {
	TokenLiteral() string
	String() string
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

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}

	return out.String()
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

func (cs *CriaStatement) String() string {
	var out bytes.Buffer

	out.WriteString(cs.TokenLiteral() + " ")
	out.WriteString(cs.Name.String())
	out.WriteString(" = ")

	if cs.Value != nil {
		out.WriteString(cs.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}
func (i *Identifier) String() string {
	return i.Value
}

type DevolveStatement struct {
	Token       token.Token
	ReturnValue Expression
}

func (ds *DevolveStatement) statementNode() {}

func (ds *DevolveStatement) TokenLiteral() string {
	return ds.Token.Literal
}

func (ds *DevolveStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ds.TokenLiteral() + " ")

	if ds.ReturnValue != nil {
		out.WriteString(ds.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

type ExpressionStatement struct {
	Token      token.Token
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}

	return ""
}

type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}
func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

type PrefixExpression struct {
	Token    token.Token
	Operator string
	Right    Expression
}

func (pe *PrefixExpression) expressionNode() {}
func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}
func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}

type InfixExpression struct {
	Token    token.Token
	Left     Expression
	Operator string
	Right    Expression
}

func (ie *InfixExpression) expressionNode() {}
func (ie *InfixExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *InfixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString(" " + ie.Operator + " ")
	out.WriteString(ie.Right.String())
	out.WriteString(")")

	return out.String()
}

type Boolean struct {
	Token token.Token
	Value bool
}

func (b *Boolean) expressionNode() {}
func (b *Boolean) TokenLiteral() string {
	return b.Token.Literal
}
func (b *Boolean) String() string {
	return b.Token.Literal
}

type SeExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (se *SeExpression) expressionNode() {}
func (se *SeExpression) TokenLiteral() string {
	return se.Token.Literal
}
func (se *SeExpression) String() string {
	var out bytes.Buffer

	out.WriteString("se")
	out.WriteString(se.Condition.String())
	out.WriteString(" ")
	out.WriteString(se.Consequence.String())

	if se.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(se.Alternative.String())
	}

	return out.String()
}

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) statementNode() {}
func (bs *BlockStatement) TokenLiteral() string {
	return bs.Token.Literal
}
func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

type ProcLiteral struct {
	Token      token.Token
	Parameters []*Identifier
	Body       *BlockStatement
}

func (pl *ProcLiteral) expressionNode() {}
func (pl *ProcLiteral) TokenLiteral() string {
	return pl.Token.Literal
}
func (pl *ProcLiteral) String() string {
	var out bytes.Buffer

	params := []string{}

	for _, p := range pl.Parameters {
		params = append(params, p.String())
	}

	out.WriteString(pl.TokenLiteral())
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") ")
	out.WriteString(pl.Body.String())

	return out.String()
}

type CallExpression struct {
	Token     token.Token
	Function  Expression
	Arguments []Expression
}

func (ce *CallExpression) expressionNode() {}
func (ce *CallExpression) TokenLiteral() string {
	return ce.Token.Literal
}
func (ce *CallExpression) String() string {
	var out bytes.Buffer

	args := []string{}

	for _, a := range ce.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(ce.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}

type StringLiteral struct {
	Token token.Token
	Value string
}

func (sl *StringLiteral) expressionNode() {}
func (sl *StringLiteral) TokenLiteral() string {
	return sl.Token.Literal
}

func (sl *StringLiteral) String() string {
	return sl.Token.Literal
}

type ListaLiteral struct {
	Token    token.Token
	Elements []Expression
}

func (ll *ListaLiteral) expressionNode() {}
func (ll *ListaLiteral) TokenLiteral() string {
	return ll.Token.Literal
}
func (ll *ListaLiteral) String() string {
	var out bytes.Buffer
	elements := []string{}

	for _, ll := range ll.Elements {
		elements = append(elements, ll.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

type IndexExpression struct {
	Token token.Token
	Left  Expression
	Index Expression
}

func (ie *IndexExpression) expressionNode() {}
func (ie *IndexExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IndexExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(ie.Left.String())
	out.WriteString("[")
	out.WriteString(ie.Index.String())
	out.WriteString("])")
	return out.String()
}

type RepeteExpression struct {
	Token     token.Token
	From      CriaStatement
	Condition Expression
	Body      BlockStatement
}

func (rp *RepeteExpression) expressionNode() {}
func (rp *RepeteExpression) TokenLiteral() string {
	return rp.Token.Literal
}
func (rp *RepeteExpression) String() string {
	var out bytes.Buffer

	out.WriteString("Repete")
	out.WriteString("( ")
	out.WriteString(rp.From.String())
	out.WriteString("ate")
	out.WriteString(rp.Condition.String())
	out.WriteString(")")
	out.WriteString(rp.Body.String())

	return out.String()

}
