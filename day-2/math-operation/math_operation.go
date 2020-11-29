package main

import (
	"fmt"
)

func main() {
	var a byte = 10
	var b byte = 20
	var c byte = a + b

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)

	a++
	b--
	d := -c // d --> byte --> 256 | 256 - c = 226

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(d)
}
