package evaluator

import (
	"fmt"

	"github.com/itkq/go-minruby-interpreter/object"
)

var builtins = map[string]*object.Builtin{
	"p": &object.Builtin{
		Fn: func(args ...object.Object) object.Object {
			for _, arg := range args {
				fmt.Println(arg.Inspect())
			}

			return NULL
		},
	},
}
