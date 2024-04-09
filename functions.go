package main

import (
	"fmt"
)

func swap(x, y string) (string, string) {

	fmt.Println("Hello, World!")
	return y, x
}
func split(x, y int) (a, b int) {
	a = x / 2
	b = y / 2
	return x, y
}
