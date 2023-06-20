package main

import "fmt"

func main() {
	anonascimento := 1996
	anoatual := 2023

	fmt.Printf("Eu tenho %v anos. \nE JÁ PASSEI POR TODOS ESSES ANOS: \n", (anoatual - anonascimento))

	for i := anonascimento; i <= 2023; i++ {
		fmt.Printf("- %v ", i)
	}
}
