package main

import "fmt"

type cliente struct {
	nome      string
	sobrenome string
	fumante   bool
}

func main() {
	cliente1 := cliente{
		nome:      "Mateus",
		sobrenome: "Bispo",
		fumante:   false,
	}
	cliente2 := cliente{"Fernando", "Santos", true}

	fmt.Println(cliente1)
	fmt.Println(cliente2)
}
