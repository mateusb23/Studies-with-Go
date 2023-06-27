package main

import "fmt"

type pessoa struct {
	nome           string
	sobrenome      string
	saboressorvete []string
}

func main() {
	pessoa1 := pessoa{
		nome:           "Mateus",
		sobrenome:      "Bispo",
		saboressorvete: []string{"Morango", "Limão", "Coco"},
	}
	pessoa2 := pessoa{
		nome:           "Bruna",
		sobrenome:      "Bilba",
		saboressorvete: []string{"Chocolate", "Doce de Leite", "Graviola"},
	}

	fmt.Println(pessoa1.nome, pessoa1.sobrenome)
	for _, v := range pessoa1.saboressorvete {
		fmt.Println(v)
	}
	fmt.Println(pessoa2.nome, pessoa2.sobrenome)
	for _, v := range pessoa2.saboressorvete {
		fmt.Println(v)
	}
}
