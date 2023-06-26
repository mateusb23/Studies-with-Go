package main

import "fmt"

func main() {
	slice1 := []int{34, 556, 77, 100, 239, 92, 21, 62, 42, 20}
	fmt.Println(slice1[:3])
	fmt.Println(slice1[4:])
	fmt.Println(slice1[2:9])
	fmt.Println(slice1[len(slice1)-2])
}
