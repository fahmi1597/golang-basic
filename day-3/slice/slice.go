package main

import (
	"fmt"
)

func main() {
	
	// both of these are array
	arr1 := [3]int{1, 2, 3}
	var arr2 = [...]int{1, 2, 3}
	fmt.Println(arr1)
	fmt.Println(arr2)
	
	// slice
	slice := []int{1, 2, 3}
	fmt.Println(slice)

}
