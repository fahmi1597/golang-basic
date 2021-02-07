package main

import "fmt"

func main() {
	matriks := [2][3]string{
		{"1,1", "1,2"},
		{"2,1", "2,2", "2,3"},
	}

	fmt.Println(matriks)
	fmt.Println("Get sub element 1,1 : ", matriks[0][0])
	fmt.Println("Get sub element 2,3 : ", matriks[1][2])

}
