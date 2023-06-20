package main

import (
	"fmt"
)

var y = 150

func main() {
	z := 20
	qualquercoisa(z)
}

func qualquercoisa(x int) {
	fmt.Println(x)
	fmt.Println(y)
}
