package main

import "fmt"

var operation int
var par int
var impar int

func main() {
	for i := 50; i <= 90; i++ {
		operation = i % 2
		if operation == 0 {
			fmt.Println("Número PAR")
			par++
		} else {
			fmt.Println("Número ÍMPAR")
			impar++
		}
	}

	fmt.Printf("\n A sequência numérica a existem %v números pares\ne %v números ímpares", par, impar)
}
