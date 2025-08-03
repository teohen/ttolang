package loader

import (
	"fmt"
	"log"
	"os"

	"github.com/teohen/ttolang/evaluator"
	"github.com/teohen/ttolang/lexer"
	"github.com/teohen/ttolang/object"
	"github.com/teohen/ttolang/parser"
	"github.com/teohen/ttolang/repl"
	"github.com/teohen/ttolang/utils"
)

func Load(path string) {
	env := object.NewEnvironment()
	err, code := utils.LoadFile("./" + path)

	if err != nil {
		log.Fatalf("Error opening .tto file in path (%s): %s", path, err)
	}

	l := lexer.New(code)
	p := parser.New(l)

	program := p.ParseProgram()

	if len(p.Errors()) != 0 {
		repl.PrintParserErrors(os.Stdout, p.Errors())
		return
	}

	evaluated := evaluator.Eval(program, env)

	if evaluated != nil {
		fmt.Println(evaluated.Inspect())
	}
}
