package main

import (
	"fmt"
)

var numero1 int
var numero2 float64
var cidade string

func main() {
	x := 10
	y := "Olá pessoal"
	z := 10.0

	fmt.Println(x, y)
	fmt.Printf("x: %v, %T\n", x, x)
	fmt.Printf("y: %v, %T\n", y, y)
	fmt.Printf("z: %v, %T\n\n", z, z)

	numero1 = 47
	numero2 = 133.58
	fmt.Println(numero1, numero2)
	fmt.Printf("numero 1: %v, %T\n", numero1, numero1)
	fmt.Printf("numero 2: %v, %T\n\n", numero2, numero2)
	
	cidade = "Recife"
	fmt.Printf("Cidade: %v, %T\n", cidade, cidade)
	fmt.Println(cidade, "\n")

	/* alguns sinais de comparação */
	numero3 := 10 == 10
	fmt.Println(numero3)

	numero4 := 10 > 20
	fmt.Println(numero4)

	numero5 := 10 == 20
	fmt.Println(numero5)

	numero6 := 10 == 10 && 10 < 20
	fmt.Println(numero6)
}
