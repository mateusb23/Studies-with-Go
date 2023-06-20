package main

import "fmt"

func main() {
	x := 10

	switch {
	case x < 5:
		fmt.Println("x é menor que 5")
	case x == 5:
		fmt.Println("x é igual a 5")
	case x > 5:
		fmt.Println("x é maior que 5")
	default:
		fmt.Println("Deu a ua louca no Go kkkk")
	}
}
