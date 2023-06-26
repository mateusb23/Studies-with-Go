package main

import "fmt"

func main() {
	// Map
	amigos := map[string]int{
		"Alfredo": 5551234,
		"Joana":   9997861,
	}
	fmt.Println(amigos)
	fmt.Println(amigos["Joana"])

	// Adicionando novos objetos ao Map
	amigos["Fernanda"] = 4459032
	fmt.Println(amigos)
}
