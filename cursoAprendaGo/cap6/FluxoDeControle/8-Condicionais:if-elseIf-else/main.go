package main

import "fmt"

func main() {
	if x := 10; x > 100 {
		fmt.Println("x É MAIOR QUE 100")
	} else if x < 10 {
		fmt.Println("x NÃO É MENOR QUE 10")
	} else {
		fmt.Println("x NÃO É MENOR QUE 10 E NEM MAIOR QUE 100")
	}
}
