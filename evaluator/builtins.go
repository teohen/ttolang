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
			case *object.String:
				return &object.Integer{Value: int64(len(arg.Value))}
			default:
				return newError("tipo de parametro de 'tam' errado. recebeu %s", args[0].Type())
			}
		},
	},
}
