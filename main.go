package main

import (
	"fmt"
	"os"

	"github.com/teohen/ttolang/repl"
)

func main() {
	fmt.Printf("Essa é a linguagem tto! Fique a vontade para digitar comandos\n")
	repl.Start(os.Stdin, os.Stdout)
}
