package main

type Pessoa struct {
	Nome      string
	Sobrenome string
	Idade     string
	Status    bool
	cpf       string
}

type Categoria struct {
	Nome string
	Pai  *Categoria
}

func main() {

}
