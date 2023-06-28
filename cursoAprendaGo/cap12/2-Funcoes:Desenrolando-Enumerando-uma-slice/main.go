package main

import "fmt"

func main() {
	si := []int{10, 10, 1, 5, 54}
	total := retornaSoma(si...)
	fmt.Println(total)
}

func retornaSoma(x ...int) int {
	y := 0
	for _, v := range x {
		y += v
	}
	return y
}
