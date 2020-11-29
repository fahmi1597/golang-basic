package main

import "fmt"

func main() {
	for a := 0; a < 3; a++ {
		if a == 2 {
			break
		}
		fmt.Println("a = ", a)
	}

	for i := 0; i < 10; i++ {
		if i%2 != 0 {
			fmt.Println("Ganjil ", i)
		}
		continue
	}
	for j := 0; j < 10; j++ {
		if j%2 == 0 {
			fmt.Println("Genap ", j)
		}
		continue
	}
}
