package main

import (
	"fmt"
)

var j string

func main() {
	for i := 33; i <= 122; i++ {
		j = string(i)
		fmt.Printf("Decimal: %d \t- Hexadecimal: %#x \t- Unicode: %#U \t- String: %q \n", i, i, i, j)
	}
}
