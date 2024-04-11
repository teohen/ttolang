package evaluator

import (
	"fmt"

	"github.com/teohen/ttolang/ast"
	"github.com/teohen/ttolang/object"
)

var (
	TRUE  = &object.Boolean{Value: "vdd"}
	FALSE = &object.Boolean{Value: "falso"}
	NULL  = &object.Null{}
)

func Eval(node ast.Node, env *object.Environment) object.Object {
	switch node := node.(type) {

	case *ast.Program:
		return evalProgram(node.Statements, env)
	case *ast.ExpressionStatement:
		return Eval(node.Expression, env)
	case *ast.PrefixExpression:
		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}
		return evalPrefixExpression(node.Operator, right)
	case *ast.InfixExpression:
		left := Eval(node.Left, env)

		if isError(left) {
			return left
		}

		right := Eval(node.Right, env)
		if isError(right) {
			return right
		}

		return evalInfixExpression(node.Operator, left, right)
	case *ast.BlockStatement:
		return evalBlockStatement(node, env)
	case *ast.SeExpression:
		return evalSeExpression(node, env)
	case *ast.DevolveStatement:
		val := Eval(node.ReturnValue, env)
		if isError(val) {
			return val
		}
		return &object.DevolveValue{Value: val}

	case *ast.IntegerLiteral:
		return &object.Integer{Value: node.Value}
	case *ast.Boolean:
		return nativeBoolToBooleanObject(node.Value)
	case *ast.CriaStatement:
		val := Eval(node.Value, env)
		if isError(val) {
			return val
		}

		env.Set(node.Name.Value, val)
	case *ast.Identifier:
		return evalIdentifier(node, env)
	case *ast.ProcLiteral:
		params := node.Parameters
		body := node.Body
		return &object.Proc{Parameters: params, Env: env, Body: body}
	case *ast.CallExpression:
		proc := Eval(node.Function, env)
		if isError(proc) {
			return proc
		}

		args := evalExpressions(node.Arguments, env)
		if len(args) == 1 && isError(args[0]) {
			return args[0]
		}

		return applyProc(proc, args)

	case *ast.ListaLiteral:
		elements := evalExpressions(node.Elements, env)
		if len(elements) == 1 && isError(elements[0]) {
			return elements[0]
		}

		return &object.Lista{Elements: elements}
	case *ast.IndexExpression:
		left := Eval(node.Left, env)

		if isError(left) {
			return left
		}

		index := Eval(node.Index, env)
		if isError(index) {
			return index
		}

		return evalIndexExpression(left, index)

	case *ast.StringLiteral:
		return &object.String{Value: node.Value}

	case *ast.AssignStatement:

		val := Eval(node.AssignExpression, env)
		if isError(val) {
			return val
		}

		env.Set(node.Name.Value, val)
		return val

	case *ast.RepeteExpression:
		evalRepeteExpression(node.Step, node.Condition, node.RepeatingStatements, env)
	}

	return nil
}

func nativeBoolToBooleanObject(input bool) *object.Boolean {
	if input {
		return TRUE
	}
	return FALSE
}

func evalProgram(stmts []ast.Statement, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range stmts {
		result = Eval(statement, env)

		switch result := result.(type) {
		case *object.DevolveValue:
			return result.Value
		case *object.Error:
			return result
		}
	}

	return result
}

func evalPrefixExpression(operator string, right object.Object) object.Object {
	switch operator {
	case "!":
		return evalBangOperatorExpression(right)
	case "-":
		return evalMinusPrefixOperatorExpression(right)

	default:
		return newError("operador desconhecido: %s %s", operator, right.Type())
	}
}

func evalBangOperatorExpression(right object.Object) object.Object {
	switch right {
	case TRUE:
		return FALSE
	case FALSE:
		return TRUE
	case NULL:
		return TRUE
	default:
		return FALSE
	}

}

func evalMinusPrefixOperatorExpression(right object.Object) object.Object {
	if right.Type() != object.INTEGER_OBJ {
		return newError("operador desconhecido: -%s", right.Type())
	}

	value := right.(*object.Integer).Value
	return &object.Integer{Value: -value}
}

func evalInfixExpression(operator string, left, right object.Object) object.Object {
	switch {
	case left.Type() == object.INTEGER_OBJ && right.Type() == object.INTEGER_OBJ:
		return evalIntegerInfixExpression(operator, left, right)
	case operator == "=":
		return nativeBoolToBooleanObject(left == right)
	case operator == "!=":
		return nativeBoolToBooleanObject(left != right)
	case left.Type() != right.Type():
		return newError("tipos diferentes: %s %s %s", left.Type(), operator, right.Type())
	case left.Type() == object.STRING_OBJ && right.Type() == object.STRING_OBJ:
		return evalStringInfixExpression(operator, left, right)
	default:
		return newError("operador desconhecido: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalIntegerInfixExpression(operator string, left, right object.Object) object.Object {
	leftVal := left.(*object.Integer).Value
	rightVal := right.(*object.Integer).Value

	switch operator {
	case "+":
		return &object.Integer{Value: leftVal + rightVal}
	case "-":
		return &object.Integer{Value: leftVal - rightVal}
	case "*":
		return &object.Integer{Value: leftVal * rightVal}
	case "/":
		return &object.Integer{Value: leftVal / rightVal}

	case "<":
		return nativeBoolToBooleanObject(leftVal < rightVal)
	case ">":
		return nativeBoolToBooleanObject(leftVal > rightVal)
	case "=":
		return nativeBoolToBooleanObject(leftVal == rightVal)
	case "!=":
		return nativeBoolToBooleanObject(leftVal != rightVal)
	default:
		return newError("operador desconhecido: %s %s %s", left.Type(), operator, right.Type())
	}
}

func evalSeExpression(se *ast.SeExpression, env *object.Environment) object.Object {
	condition := Eval(se.Condition, env)

	if isError(condition) {
		return condition
	}

	if isTruthy(condition) {
		return Eval(se.Consequence, env)
	} else if se.Alternative != nil {
		return Eval(se.Alternative, env)
	} else {
		return NULL
	}
}

func isTruthy(obj object.Object) bool {
	switch obj {
	case NULL:
		return false
	case TRUE:
		return true
	case FALSE:
		return false
	default:
		return true
	}
}

func evalBlockStatement(block *ast.BlockStatement, env *object.Environment) object.Object {
	var result object.Object

	for _, statement := range block.Statements {
		result = Eval(statement, env)

		if result != nil {
			rt := result.Type()
			if rt == object.DEVOLVE_VALUE_OBJ || rt == object.ERROR_OBJ {
				return result
			}
		}
	}

	return result
}

func newError(format string, a ...interface{}) *object.Error {
	return &object.Error{Message: fmt.Sprintf(format, a...)}
}

func isError(obj object.Object) bool {
	if obj != nil {
		return obj.Type() == object.ERROR_OBJ
	}

	return false
}

func evalIdentifier(node *ast.Identifier, env *object.Environment) object.Object {
	if val, ok := env.Get(node.Value); ok {
		return val
	}

	if builtin, ok := builtins[node.Value]; ok {
		return builtin
	}

	return newError("identificador é desconhecido: " + node.Value)
}

func evalExpressions(exps []ast.Expression, env *object.Environment) []object.Object {
	var result []object.Object

	for _, e := range exps {
		evaluated := Eval(e, env)
		if isError(evaluated) {
			return []object.Object{evaluated}
		}

		result = append(result, evaluated)
	}

	return result
}

func applyProc(pc object.Object, args []object.Object) object.Object {

	switch proc := pc.(type) {
	case *object.Proc:
		extendedEnv := extendProcEnv(proc, args)

		evaluated := Eval(proc.Body, extendedEnv)

		return unwrapDevolveValue(evaluated)

	case *object.Builtin:
		return proc.Fn(args...)

	default:
		return newError("não é um proc: %s", proc.Type())
	}

}

func extendProcEnv(proc *object.Proc, args []object.Object) *object.Environment {
	env := object.NewEnclosedEnvironment(proc.Env)

	for paramIdx, param := range proc.Parameters {
		env.Set(param.Value, args[paramIdx])
	}

	return env
}

func unwrapDevolveValue(obj object.Object) object.Object {
	if devolveValue, ok := obj.(*object.DevolveValue); ok {
		return devolveValue.Value
	}

	return obj
}

func evalStringInfixExpression(operator string, left, right object.Object) object.Object {
	if operator != "+" {
		return newError("operador desconhecido: %s %s %s", left.Type(), operator, right.Type())
	}

	lefVal := left.(*object.String).Value
	rigVal := right.(*object.String).Value
	return &object.String{Value: lefVal + rigVal}
}

func evalIndexExpression(left, index object.Object) object.Object {
	switch {
	case left.Type() == object.LISTA_OBJ && index.Type() == object.INTEGER_OBJ:
		return evalListaIndexExpression(left, index)
	default:
		return newError("operador de indice não aceito: %s", left.Type())
	}
}

func evalListaIndexExpression(lista, index object.Object) object.Object {
	listaObject := lista.(*object.Lista)
	idx := index.(*object.Integer).Value

	max := int64(len(listaObject.Elements) - 1)
	// TODO: retornar um erro ao acessar um array com indice nao existente
	if idx < 0 || idx > max {
		return NULL
	}

	return listaObject.Elements[idx]
}

func evalRepeteExpression(step *ast.AssignStatement, condition ast.Expression, repeating *ast.BlockStatement, env *object.Environment) {
	shouldLoop := false

	conditional := Eval(condition, env)

	shouldLoop = !isTruthy(conditional)
	for shouldLoop {
		newEnv := object.NewEnclosedEnvironment(env)

		_ = Eval(repeating, newEnv)

		val := Eval(step, newEnv)

		newEnv.Set(step.Name.Value, val)

		conditionalNewEnv := Eval(condition, newEnv)

		shouldLoop = !isTruthy(conditionalNewEnv)
	}

}
