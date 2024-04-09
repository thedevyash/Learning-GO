package main

import (
	"fmt"
	"reflect"
)

func test() {
	var a int
	var b = "10"
	c := 20
	fmt.Printf("a=%d, b=%v, c=%d\n", a, reflect.TypeOf(b), c)
}

func main() {
	fmt.Println("Hello, World!")
	test()
}
