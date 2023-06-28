package main

import "fmt"

func main() {
	basica()
	argumento("tarde")
	fmt.Println("Soma:", retornoSoma(2, 7))
	fmt.Println("-----------------------------------")
	total, quantos, mensagem := retornosMultiplosSoma("--->  Terminei de estudar  <---", 10, 10, 1, 5, 54)
	fmt.Println(total, quantos, mensagem)
}

func basica() {
	fmt.Println("Oi, bom dia!")
}

func argumento(s string) {
	if s == "manhã" {
		fmt.Println("Oi, tenha um bom dia!")
	}
	if s == "tarde" {
		fmt.Println("Oi, tenha uma boa tarde!")
	}
	if s == "noite" {
		fmt.Println("oi, tenha uma ótima noite!")
	}
}

func retornoSoma(x, y int) int {
	return x + y
}

// func retornosMultiplosSoma(x ...int) (int, int, string) { // Esses ...int Indica que a função poderá receber uma quantidade indeterminda ou zero de valores do tipo int. E com isso, o x ali passa a ser um Slice
// 	y := 0
// 	for _, v := range x {
// 		y += v
// 	}
// 	return y, len(x), "--->  Terminei de estudar  <---"
// }

// OU

func retornosMultiplosSoma(s string, x ...int) (int, int, string) { // Nesse Caso é importante Ressaltar que o elemento variádico x ...int sempre de ser o último parâmetro
	y := 0
	for _, v := range x {
		y += v
	}
	return y, len(x), s
}
