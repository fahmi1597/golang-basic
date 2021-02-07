package main

import "fmt"

func magic(i interface{}) interface{} {
	return i
}

func main() {

	fmt.Println(magic("string"))
	fmt.Println(magic(100))
}
