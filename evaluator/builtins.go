package evaluator

import (
	"fmt"

	"github.com/teohen/ttolang/object"
)

var builtins = map[string]*object.Builtin{
	"tam": {
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
	"anexar": {
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 2 {
				return newError("quantidade errada de parametros. recebeu=%d, aceita=2", len(args))
			}

			switch args[0].(type) {
			case *object.Lista:
				arr := args[0].(*object.Lista)
				length := len(arr.Elements)

				newElements := make([]object.Object, length+1)
				copy(newElements, arr.Elements)
				newElements[length] = args[1]

				return &object.Lista{Elements: newElements}

			case *object.String:
				if args[1].Type() == object.STRING_OBJ {
					string := args[0].(*object.String)
					secondString := args[1].(*object.String)
					return &object.String{Value: string.Value + secondString.Value}
				}
				return newError("segundo parametro com tipo errado. Recebeu=%s, aceita=STRING", args[1].Type())

			default:
				return newError("primeiro parametro com tipo errado. recebeu=%s, aceita=LISTA ou STRING", args[0].Type())
			}
		},
	},
	"mostra": {
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return NULL
		},
	},
}
