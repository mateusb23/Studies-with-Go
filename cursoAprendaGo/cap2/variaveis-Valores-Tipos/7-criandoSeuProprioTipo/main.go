package main

import "fmt"

type hotdog int // criamos um tipo customizado chamado hotdog. E esse tipo hotdog tem um subtipo subjacente, que é int

var b hotdog // e aqui criamos uma variável chamada b do tipo hotdog

var c hotdog = 35

func main() {
	fmt.Printf("Tipo da variável b: %T\n", b) // aqui podemos constatar qual é o tipo da variável b, a qual declaramos na linha 7.

	b = 35
	d := 35

	fmt.Printf("Tipo da variável c: %T\n", c)
	fmt.Printf("Tipo da variável d: %T\n\n", d)

	fmt.Println("VALOR DE b:", b)
	fmt.Println("VALOR DE c:", c)
	fmt.Println("VALOR DE d:", d, "\n")

	if b == c {
		fmt.Println("b É IGUAL A c")
	} else {
		fmt.Println("b É DIFERENTE DE c")
	}

	/*

		---> Se salvarmos esse trecho de código que podemos ver logo abaixo e tentar dar um go run nele,
		veremos o Go nos mostrar a mensagem --> # command-line-arguments./main.go:27:10: invalid
		operation: b != d (mismatched types hotdog and int)  <--

		-> OU SEJA, ELE NOS DEIXA CLARO QUE NÃO TEM COMO COMPARAR DUAS VARIÁVEIS DE TIPOS
		DIFERENTES, VISTO QUE -------> b É DO TIPO hotdog <-------- E ------> c É DO TIPO int <--

		OBS: NÃO DEVEMOS CONFUNDIR tipo COM subtipo.

		if b != d {
			fmt.Println("b É DIFERENTE DE d")
		}
	*/
}
