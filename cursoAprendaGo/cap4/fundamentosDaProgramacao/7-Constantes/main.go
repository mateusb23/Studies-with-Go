package main

import "fmt"

const x = 10 // as Constantes não tipadas só terão um tipo atribuido a elas quando forem usadas.

var y = 10 // o tipo de uma Variável é definido no momento da atribuição

const (
	a = 10
	b = 20
	c = 30
)

var w int
var z float64

func main() {
	z = x
	fmt.Printf("Valor: %v\n Tipo: %T", z, z)
}
