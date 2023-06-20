package main

import (
	"fmt"
)

func main() {
	fmt.Println("Hello World!")

	x := 16
	y := "Mateus"
	z := true

	fmt.Println(x, y, z)

	numerodebytes, erros := fmt.Println("Hello World", "Oi galera", 100)
	fmt.Println(numerodebytes, erros)
}
