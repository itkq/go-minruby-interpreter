package main

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/itkq/go-minruby-interpreter/evaluator"
	"github.com/itkq/go-minruby-interpreter/lexer"
	"github.com/itkq/go-minruby-interpreter/object"
	"github.com/itkq/go-minruby-interpreter/parser"
	"github.com/itkq/go-minruby-interpreter/repl"
)

func main() {
	if len(os.Args) == 2 {
		data, err := ioutil.ReadFile(os.Args[1])
		if err != nil {
			fmt.Println("read file error")
			os.Exit(1)
		}
		body := string(data)
		l := lexer.New(body)
		p := parser.New(l)
		program := p.ParseProgram()
		env := object.NewEnvironment()
		if len(p.Errors()) != 0 {
		}
		evaluator.Eval(program, env)
	} else {
		repl.Start(os.Stdin, os.Stdout)
	}
}
