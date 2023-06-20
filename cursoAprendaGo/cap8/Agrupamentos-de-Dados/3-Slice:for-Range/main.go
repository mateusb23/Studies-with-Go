package main

import "fmt"

func main() {
	// COM A FUNÇÃO range NÓS CONSEGUIMOS PASSAR POR CADA ÍNDICE(POSIÇÃO) DO slice e
	// ATRIBUIR DETERMINADOS COMPORTAMENTOS QUE DESEJARMOS

	slice := []string{"banana", "maçã", "jaca", "uva"}

	for indice, valor := range slice {
		fmt.Println("No índice", indice, "temos o valor:", valor)
	}

	slice[3] = "melancia"

	for indice, valor := range slice {
		fmt.Println("No índice", indice, "temos o valor:", valor)
	}
}
