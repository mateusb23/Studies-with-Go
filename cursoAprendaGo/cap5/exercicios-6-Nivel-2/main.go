package main

import "fmt"

const (
	a = iota + 2023
	b
	c
	d
	e
)

func main() {
	fmt.Println(a, b, c, d, e)
}
