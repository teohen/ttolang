package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/teohen/ttolang/ast"
)

type ObjectType string

type BuiltinFunction func(args ...Object) Object

const (
	INTEGER_OBJ       = "INTEIRO"
	BOOLEAN_OBJ       = "BOOLEANO"
	NULL_OBJ          = "NULO"
	DEVOLVE_VALUE_OBJ = "DEVOLVE_VALUE"
	ERROR_OBJ         = "ERRO"
	FUNCTION_OBJ      = "PROCESSO"
	STRING_OBJ        = "STRING"
	BUILTIN_OBJ       = "BUILTIN"
	LISTA_OBJ         = "LISTA"
	ESTRUTURA_OBJ     = "ESTRUTURA"
)

type Object interface {
	Type() ObjectType
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Inspect() string {
	return fmt.Sprintf("%d", i.Value)
}

func (i *Integer) Type() ObjectType {
	return INTEGER_OBJ
}

type Boolean struct {
	Value string
}

func (b *Boolean) Type() ObjectType {
	return BOOLEAN_OBJ
}
func (b *Boolean) Inspect() string {
	return b.Value
}

type Null struct {
}

func (n *Null) Type() ObjectType {
	return NULL_OBJ
}

func (n *Null) Inspect() string {
	return "null"
}

type DevolveValue struct {
	Value Object
}

func (dv *DevolveValue) Type() ObjectType {
	return DEVOLVE_VALUE_OBJ
}

func (dv *DevolveValue) Inspect() string {
	return dv.Value.Inspect()
}

type Error struct {
	Message string
}

func (e *Error) Type() ObjectType {
	return ERROR_OBJ
}
func (e *Error) Inspect() string {
	return "Problema: " + e.Message
}

type Environment struct {
	store map[string]Object
	outer *Environment
}

func NewEnvironment() *Environment {
	s := make(map[string]Object)
	return &Environment{store: s, outer: nil}
}

func NewEnclosedEnvironment(outer *Environment) *Environment {
	env := NewEnvironment()
	env.outer = outer
	return env
}

func (e *Environment) Get(name string) (Object, bool) {
	obj, ok := e.store[name]

	if !ok && e.outer != nil {
		obj, ok = e.outer.Get(name)
	}
	return obj, ok
}

func (e *Environment) Set(name string, val Object) Object {
	if e.outer != nil {
		_, existsOnOuter := e.outer.Get(name)

		if existsOnOuter {
			e.outer.Set(name, val)
			return val
		}
	}
	e.store[name] = val

	return val
}

type Proc struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (p *Proc) Type() ObjectType {
	return FUNCTION_OBJ
}

func (p *Proc) Inspect() string {
	var out bytes.Buffer

	params := []string{}

	for _, p := range p.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("proc")
	out.WriteString("(")
	out.Write([]byte(strings.Join(params, " ,")))
	out.WriteString(") {\n")
	out.WriteString(p.Body.String())
	out.WriteString("\n}")

	return out.String()
}

type String struct {
	Value string
}

func (s *String) Type() ObjectType {
	return STRING_OBJ
}
func (s *String) Inspect() string {
	return s.Value
}

type Builtin struct {
	Fn BuiltinFunction
}

func (b *Builtin) Type() ObjectType {
	return BUILTIN_OBJ
}
func (b *Builtin) Inspect() string {
	return "builtin proc"
}

type Lista struct {
	Elements []Object
}

func (l *Lista) Type() ObjectType {

	return LISTA_OBJ
}

func (l *Lista) Inspect() string {
	var out bytes.Buffer

	elements := []string{}

	for _, e := range l.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}

type Estrutura struct {
	Items map[string]Object
}

func (t *Estrutura) Type() ObjectType {

	return ESTRUTURA_OBJ
}

func (t *Estrutura) Inspect() string {
	var out bytes.Buffer

	out.WriteString("{")
	for _, e := range t.Items {
		out.WriteString(e.Inspect())
	}
	out.WriteString("}")

	return out.String()
}
