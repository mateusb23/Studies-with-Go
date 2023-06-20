package main

import "fmt"

var x int // isso aqui é uma DECLARAÇÃO de variável

func main() {
	x = 45    // isso aqui é uma INICIALIZAÇÃO de variável
	y := 80.5 // isso aqui é uma DECLARAÇÃO + INICIALIZAÇÃO ao mesmo tempo

	fmt.Printf("Number x: %v, %T\n\n", x, x)
	fmt.Printf("Number y: %v, %T\n\n", y, y)

	x = 122 // isso aqui é uma ATRIBUIÇÃO de valor novo a uma variável
	fmt.Printf("Number x: %v, %T\n\n", x, x)

	// OBS1: SEMPRE QUE POSSÍVEL, UTILIZE :=
	// OBS2: ATRIBUIÇÕES SÓ PODEM SER FEITAS EM CODE BLOCKS (dentro de func)

	/* RESUMO:

	1) toda INICIALIZAÇÃO COM VALOR é também uma ATRIBUIÇÃO
	2) porém, uma ATRIBUIÇÃO feita após um valor já ter sido INICIALIZADO anteriormente é apenas uma ATRIBUIÇÃO
	3) Por fim, toda (DECLARAÇÃO + INICIALIZAÇÃO) feita através do :=  é também uma ATRIBUIÇÃO */

	/* VALORE ZERO de cada tipo

	- ints: 0
	- floats: 0.0
	- booleans: false
	- strings: ""
	- pointers, functions, interfaces, slices, channels, maps: nil */

	var a int
	var b float64
	var c string
	var d bool

	fmt.Printf("Value a: %v, %T\n\n", a, a)
	fmt.Printf("Value b: %v, %T\n\n", b, b)
	fmt.Printf("Value c: %v, %T\n\n", c, c)
	fmt.Printf("Value a: %v, %T\n\n", d, d)
}
