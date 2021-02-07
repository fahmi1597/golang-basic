package main

import "fmt"

func calcSum(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total = total + value
	}
	return total
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(calcSum(slice...))
}
