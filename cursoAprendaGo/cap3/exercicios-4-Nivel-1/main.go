package main

import "fmt"

type nota int

var x nota

func main() {

	fmt.Printf("Valor da variável x: %v\nTipo da variável x: %T\n", x, x)

	x = 42

	fmt.Printf("Valor Atualizado da variável x: %v", x)
}
