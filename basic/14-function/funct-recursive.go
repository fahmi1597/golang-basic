package main

import "fmt"

func faktorialLoop(nilai int) int {
	result := 1
	for i := nilai; i > 0; i-- {
		result = result * i
	}
	return result
}
func faktorialRecursive(nilai int) int {
	if nilai == 1 {
		return 1
	} else {
		return n * faktorialRecursive(nilai-1)
	}
}

func main() {
	fmt.Println(faktorialLoop(5))
	fmt.Println(faktorialRecursive(5))

}
