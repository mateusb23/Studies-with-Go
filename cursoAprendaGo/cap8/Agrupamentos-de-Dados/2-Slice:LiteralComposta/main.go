package main

import "fmt"

func main() {
	// Array sempre vai ter um tamanho Finito e fixo
	array := [7]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	fmt.Println(len(array))

	// Slice pode crescer, sempre que desejarmos e ele precisar. Ele não tem um tamanho fixo
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	fmt.Println(slice)
	fmt.Println(len(slice))
}
