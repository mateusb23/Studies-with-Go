package main

import "fmt"

func main() {
	umaslice := []int{1, 2, 3, 4}
	outraslice := []int{9, 10, 11, 12}

	fmt.Println(umaslice)

	umaslice = append(umaslice, 5, 6, 7, 8)
	fmt.Println(umaslice)

	umaslice = append(umaslice, outraslice...) // Esses 3 potinhos(...) está dizendo que é para pegar os índices(valores) que estão dentro desse Slice
	fmt.Println(umaslice)
}
