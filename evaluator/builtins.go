package evaluator

import (
	"fmt"

	"github.com/itkq/go-minruby-interpreter/object"
)

var builtins = map[string]*object.Builtin{
	"p": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			if len(args) != 1 {
				return newError("wrong number of arguments. got=%d, want=1", len(args))
			}

			return args[0] // XXX:
		},
	},
}
