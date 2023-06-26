package main

import "fmt"

func main() {
	slice1 := []int{34, 556, 77, 100, 239, 92, 21, 62, 42, 20}

	for _, v := range slice1 {
		fmt.Println(v)
	}

	fmt.Printf("Tipo do Slice: %T", slice1)

}
