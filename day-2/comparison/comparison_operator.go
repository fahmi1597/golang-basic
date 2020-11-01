package main

import (
	"fmt"
)

func main() {
	var (
		a = 100
		b = 101
		c = "Fahmi"
		d = "J"
	)
	fmt.Println(a > b)
	fmt.Println(a < b)
	fmt.Println(a == b)
	fmt.Println(a != b)

	//membandingkan nilai ascii
	fmt.Println("Compare string")
	fmt.Println(c > d)
	fmt.Println(c < d)
	fmt.Println(c == d)
	fmt.Println(c != d)
}
