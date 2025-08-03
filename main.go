package main

import (
	"fmt"
	"os"

	"github.com/teohen/ttolang/loader"
	"github.com/teohen/ttolang/repl"
)

func main() {

	args := os.Args

	if len(args) > 1 && args[1] != "" {
		loader.Load(args[1])
	} else {
		fmt.Printf("Essa é a linguagem tto! Fique a vontade para digitar comandos\n")

		repl.Start(os.Stdin, os.Stdout)
	}
}

// TODO: error when trying to redefine a existent indentifier
