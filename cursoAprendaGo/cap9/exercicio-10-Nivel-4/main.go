package main

import "fmt"

func main() {
	doguinhos := map[string]string{
		"Velhinha_Mel":       "Dormir em paz",
		"Dostoiévski_Fiódor": "Correr pra caramba",
		"Gordinha_Safira":    "Comer pipoca",
	}
	doguinhos["Capivara_Barbuda"] = "Ser acariciada"

	delete(doguinhos, "Gordinha_Safira")

	for k, v := range doguinhos {
		fmt.Println(k + ": " + v)
	}
}
