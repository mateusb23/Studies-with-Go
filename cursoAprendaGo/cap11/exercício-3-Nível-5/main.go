package main

import "fmt"

type veiculo struct {
	portas int
	cor    string
}

type caminhonete struct {
	veiculo
	tracaoQuatroRodas bool
}

type sedan struct {
	veiculo
	modeloLuxo bool
}

func main() {

	civic := sedan{
		veiculo:    veiculo{4, "preta"},
		modeloLuxo: false,
	}

	hylux := caminhonete{
		veiculo: veiculo{
			portas: 4,
			cor:    "vermelha",
		},
		tracaoQuatroRodas: true,
	}

	fmt.Println("\t", civic)
	fmt.Println("\t", hylux)
}
