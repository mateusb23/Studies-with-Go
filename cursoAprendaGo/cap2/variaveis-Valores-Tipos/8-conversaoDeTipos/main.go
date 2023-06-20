package main

import "fmt"

type hotdog int

var b hotdog = 10

func main() {
	x := 10
	fmt.Printf("Tipo de x: %T\n, Valor de x: %v\n", x, x)
	fmt.Printf("Tipo de b: %T\n ,Valor de b: %v\n", b, b)

	x = int(b) /* isso aqui é o que costumamos chamar de CONVERSÃO em Go, pois convertemos a variável
	b para o tipo int e atribuimos ela à variável x do tipo int também. */
	fmt.Printf("Tipo atualizado de x:%T\n, Valor atualizado de x: %v\n", x, x)
}
