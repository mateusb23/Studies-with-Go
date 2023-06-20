package main

import "fmt"

var x interface{}

func main() {
	x = "Mateus"

	switch x.(type) {
	case int:
		fmt.Println("É um int")
	case bool:
		fmt.Println("É um bool")
	case string:
		fmt.Println("É um string")
	case float64:
		fmt.Println("É um float")
	}
}
