package main

import "fmt"

type nota int

var x nota
var y int

func main() {

	fmt.Printf("Valor da variável x: %v\nTipo da variável x: %T\n", x, x)

	x = 42

	fmt.Printf("Valor Atualizado da variável x: %v\n", x)

	y = int(x)

	fmt.Printf("Valor da variável y: %v\nTipo da variável y: %T: ", y, y)
}
