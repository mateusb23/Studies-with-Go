package main

import "fmt"

func main() {
	slice1 := make([]int, 5, 10) // cria um slice com 5 valores zero e 10 índices de capacidade
	for i := 0; i < 4; i++ {
		slice1[i] = i
	}
	fmt.Println("Slice 1 impresso: ", slice1, "\nTamanho ocupado do slice 1: ", len(slice1), "\nCapacidade do Slice 1: ", cap(slice1))

	// Como já estabelecemos 5 valores sendo criados em um slice com capacidade 10, não temos nenhum valor inserido
	// nas outras 5 posições restantes.
	// E é justamente por isso que não temos a possibilidade de setar com slice1[5] = 35, por exemplo, pois sempre vai dar erro.
	// é para resolvermos isso utilizamos o append(), conforme fiz abaixo:

	slice1 = append(slice1, 37)
	fmt.Println("Slice 1 impresso APÓS adicionarmos mais 1 índice com valor: ", slice1)
}
