package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

type profissional struct {
	pessoa
	titulo  string
	salario int
}

func main() {
	pessoa1 := pessoa{
		nome:  "Rebeca",
		idade: 32,
	}

	trabalhador1 := profissional{
		pessoa:  pessoa1,
		titulo:  "Contabilista",
		salario: 3000,
	}

	trabalhador2 := profissional{
		pessoa: pessoa{
			nome:  "Fernanda",
			idade: 53,
		},
		titulo:  "Doceira",
		salario: 4200,
	}

	fmt.Println(trabalhador1)
	fmt.Println(trabalhador2)
}
