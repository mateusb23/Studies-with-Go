package main

import "fmt"

func main() {
	x := 0
	y := 5
	for x < y {
		fmt.Printf("%v é menor que %v\n", x, y)
		x++
	}

	fmt.Println("\n\n")

	z := 0
	for {
		if z < 10 {
			fmt.Println("x é menor que 10")
			z++
		} else {
			fmt.Println("Encerre, pois z chegou ao valor 10!!!")
			break
		}
	}
}
