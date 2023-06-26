package main

import "fmt"

func main() {
	doguinhos := [][]string{
		[]string{
			"Mel",
			"Velhinha",
			"Dormir em paz",
		},
		[]string{
			"Fiódor",
			"Dostoiévski",
			"Correr pra caramba",
		},
		[]string{
			"Safira",
			"Gordinha",
			"Comer pipoca",
		},
	}
	for _, v := range doguinhos {
		fmt.Println(v)
	}
}
