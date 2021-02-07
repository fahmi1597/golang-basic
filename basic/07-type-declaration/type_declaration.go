package main

import "fmt"

type (
	person struct {
		name string
		age  int
	}
)

type p person

func main() {
	type npm string

	var noInduk npm = "52416483"
	fmt.Println(noInduk)

}
