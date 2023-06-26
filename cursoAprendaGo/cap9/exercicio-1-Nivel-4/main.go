package main

import "fmt"

func main() {
	array1 := [5]int{144, 26, 78, 335, 50}

	for _, valor := range array1 {
		fmt.Println(valor)
	}

	fmt.Printf("Tipo do Array: %T", array1)
}
