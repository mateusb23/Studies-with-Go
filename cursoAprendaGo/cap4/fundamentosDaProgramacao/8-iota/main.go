package main

import "fmt"

const (
	a = iota * 10
	_
	c
	x
	_ = iota
	z = iota
)

func main() {
	fmt.Println(a, c, x, z)
}
