package main

import "fmt"

func main() {
	n := 200
	fmt.Printf("%d, %b, %#x\n", n, n, n)
	y := n << 1
	fmt.Printf("%d, %b, %#x", y, y, y)
}
