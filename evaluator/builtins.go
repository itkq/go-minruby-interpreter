package evaluator

import (
	"fmt"
	"strconv"

	"github.com/itkq/go-minruby-interpreter/object"
)

var builtins = map[string]*object.Builtin{
	"p": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}
			fmt.Println(args[0].Inspect())
			return args[0]
		},
	},
	"Integer": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			switch arg := args[0].(type) {
			case *object.String:
				value, err := strconv.ParseInt(arg.Value, 0, 64)
				if err != nil {
					return newError("could not parse %q as integer", arg.Value)
				}
				return &object.Integer{Value: value}
			default:
				return newError("argument to `Integer` not supported, got %s", arg.Type())
			}
		},
	},
}
