package main

import "fmt"

type pessoa struct {
	nome  string
	idade int
}

func (p pessoa) oibomdia() {
	fmt.Println(p.nome, "disse 'bom dia'!")
}

func main() {
	pessoa1 := pessoa{"Safira", 11}
	pessoa1.oibomdia()
}

/* LEMBRANDO COMO FUNCIONAM MÉTODOS E FUNÇÕES EM GO:

func (receiver) identifier(parameters) (returns) {
	code
}

----> EXEMPLO DE FUNÇÃO:

func oibomdia(p pessoa) {
	fmt.Println(p.nome, "disse 'bom dia'!")
}

----> EXEMPLO DE MÉTODO:

func (p pessoa) oibomdia() {
	fmt.Println(p.nome, "disse 'bom dia'!")
}

*/
