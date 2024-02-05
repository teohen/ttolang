package loader

import (
	"bufio"
	"fmt"
	"log"
	"os"

	"github.com/teohen/ttolang/evaluator"
	"github.com/teohen/ttolang/lexer"
	"github.com/teohen/ttolang/object"
	"github.com/teohen/ttolang/parser"
	"github.com/teohen/ttolang/repl"
)

func Load(path string) {
	file, err := os.Open("./" + path)

	if err != nil {
		panic(err)
	}

	defer func() {
		if err = file.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(file)
	env := object.NewEnvironment()
	code := ""

	for scanner.Scan() {
		code += scanner.Text()
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
