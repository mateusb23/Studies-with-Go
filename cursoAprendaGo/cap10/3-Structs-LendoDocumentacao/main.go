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

	pessoa3 := pessoa{"Mateus", 26}

	trabalhador3 := profissional{pessoa{"Bruna", 23}, "Enfermeira", 3500}

	fmt.Println(trabalhador1.pessoa.nome)
	// OU
	fmt.Println(trabalhador1.nome)
	fmt.Println(trabalhador2.pessoa)
	fmt.Println(trabalhador2)
	fmt.Println(pessoa3)
	fmt.Println(trabalhador3)
}
