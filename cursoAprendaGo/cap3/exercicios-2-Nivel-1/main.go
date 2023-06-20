package main

import "fmt"

var x int
var y string
var z bool

func main() {
	fmt.Printf("Valor de x: %v\n Tipo de x: %T\n", x, x)
	fmt.Printf("Valor de y: %v\n Tipo de y: %T\n", y, y)
	fmt.Printf("Valor de z: %v\n Tipo de z: %T\n", z, z)

	/* OBS: os valores printados na tela são o VALOR ZERO de cada variável, visto que apenas
	declaramos as variáveis x, y, z , mas não atribuimos nenhum valor à elas. */
}
