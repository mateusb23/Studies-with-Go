package main

import "fmt"

func main() {
	x := 3
	if (x == 2) || (x == 3) {
		fmt.Println("x É IGUAL A 2 OU 3")
	}
	if (x < 4) && (x > 2) {
		fmt.Println("x É EXATAMENTE IGUAL A 3")
	}
	if x%2 != 0 {
		fmt.Println("x É ÍMPAR")
	} else {
		fmt.Println("x É PAR")
	}
}
