package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     int
}

type dentista struct {
	pessoa
	dentesarrancados int
	salario          float64
}

type arquiteto struct {
	pessoa
	tipodeconstrucao string
	tamanhodaloucura string
}

func (p dentista) oibomdia() {
	fmt.Println("---> Meu nome é", p.pessoa.nome, "e eu já arranquei", p.dentesarrancados, "dentes, e ouve só: Bom dia!")
}

func (p arquiteto) oibomdia() {
	fmt.Println("---> Meu nome é", p.nome, "e ouve só: Bom dia!")
}

type gente interface {
	oibomdia()
}

func serhumano(g gente) {
	g.oibomdia()
	switch g.(type) {
	case dentista:
		fmt.Println("Eu ganho:", g.(dentista).salario)
	case arquiteto:
		fmt.Println("Eu construo:", g.(arquiteto).tipodeconstrucao)
	}
}

func main() {
	mrdente := dentista{
		pessoa: pessoa{
			nome:      "Douglas",
			sobrenome: "Beltrão",
			idade:     34,
		},
		dentesarrancados: 2139,
		salario:          4150.90,
	}

	var mrpredio arquiteto
	mrpredio = arquiteto{
		pessoa: pessoa{
			nome:      "Roberta",
			sobrenome: "Peixoto",
			idade:     41,
		},
		tipodeconstrucao: "apartamento duplex",
		tamanhodaloucura: "Muito grande",
	}

	mrdente.oibomdia()
	mrpredio.oibomdia()

	fmt.Println("----------------------------------------------------------------------")
	fmt.Println("----------------------------------------------------------------------")

	serhumano(mrdente)
	serhumano(mrpredio)
}
