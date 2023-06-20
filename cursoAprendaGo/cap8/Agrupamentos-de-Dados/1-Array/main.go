package main

import "fmt"

var x [5]int
var y [6]int
var z = [4]string{"a", "b", "c", "d"}

func main() {
	x[0] = 10
	x[1] = 45
	x[2] = 8
	w := [3]float64{2.3, 3.5, 5.1}
	fmt.Println(x[0], x[1], x[2])
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
	fmt.Println(w)
	fmt.Printf("Tipo da variável x: %T", x)
	fmt.Printf("\nTipo da variável y: %T\n", y)
	// OBS: É IMPOTANTE RESSALTAR QUE, MESMO x E y SENDO ARRAYS DE int, OS DOIS SÃO CONSIDERADOS TIPOS DIFERENTES

	// A função len(nomeDoArray) nos retorna o tamanho do Array
	fmt.Println("Tamanho do Array x =", len(x))
}
