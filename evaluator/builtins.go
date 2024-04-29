package evaluator

import (
	"fmt"

	"maps"

	"github.com/teohen/ttolang/object"
)

func validateLength(expectedLength, actualLength int) *object.Error {
	if actualLength != expectedLength {
		return newError("quantidade errada de parametros. recebeu=%d, aceita=%d", actualLength, expectedLength)
	}
	return nil
}

var builtins = map[string]*object.Builtin{
	"tam": {
		Fn: func(args ...object.Object) object.Object {
			if error := validateLength(1, len(args)); error != nil {
				return error
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
			switch args[0].(type) {
			case *object.Lista:
				if error := validateLength(2, len(args)); error != nil {
					return error
				}

				arr := args[0].(*object.Lista)
				length := len(arr.Elements)

				newElements := make([]object.Object, length+1)
				copy(newElements, arr.Elements)
				newElements[length] = args[1]

				return &object.Lista{Elements: newElements}

			case *object.String:
				if error := validateLength(2, len(args)); error != nil {
					return error
				}

				if args[1].Type() == object.STRING_OBJ {
					string := args[0].(*object.String)
					secondString := args[1].(*object.String)
					return &object.String{Value: string.Value + secondString.Value}
				}
				return newError("segundo parametro com tipo errado. Recebeu=%s, aceita=STRING", args[1].Type())

			case *object.Estrutura:
				if error := validateLength(3, len(args)); error != nil {
					return error
				}

				estr := args[0].(*object.Estrutura)
				newItems := make(map[string]object.Object)
				maps.Copy(newItems, estr.Items)

				estrParam, ok := args[1].(*object.String)
				if ok {
					newItems[estrParam.Value] = args[2]
					return &object.Estrutura{Items: newItems}
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
