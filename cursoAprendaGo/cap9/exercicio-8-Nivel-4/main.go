package main

import "fmt"

func main() {
	doguinhos := map[string]string{
		"Velhinha_Mel":       "Dormir em paz",
		"Dostoiévski_Fiódor": "Correr pra caramba",
		"Gordinha_Safira":    "Comer pipoca",
	}
	for k, v := range doguinhos {
		fmt.Println(k + ": " + v)
	}
}
