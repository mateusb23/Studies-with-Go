package main

import "fmt"

func main() {
	fmt.Println("1")
	defer fmt.Println("2") // A tradição literal da palavra "defer" é "adiar". Onde o defer estiver, sempre será a última ação a ser executada
	fmt.Println("3")
	fmt.Println("4")
	fmt.Println("5")
}
