package evaluator

import (
	"github.com/teohen/ttolang/object"
)

var builtins = map[string]*object.Builtin{
	"tam": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("quantidade errada de parametros. recebeu=%d, aceita=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.Lista:
				return &object.Integer{Value: int64(len(arg.Elements))}
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("tipo de parametro de 'tam' errado. recebeu %s", args[0].Type())
			}
		},
	},
	"anexar": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("quantidade errada de parametros. recebeu=%d, aceita=2", len(args))
			}

			if args[0].Type() != object.LISTA_OBJ {
				return newError("primeiro parametro com tipo errado. recebeu=%s, aceita=Lista", args[0].Type())
			}

			arr := args[0].(*object.Lista)
			lenght := len(arr.Elements)

			newElements := make([]object.Object, lenght+1)
			copy(newElements, arr.Elements)
			newElements[lenght] = args[1]

			return &object.Lista{Elements: newElements}
		},
	},
}
