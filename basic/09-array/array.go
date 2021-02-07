package main

import (
	"fmt"
)

func main() {

	fmt.Println("Cara 1")
	// deklarasi 1s
	var arrNama [3]string
	arrNama[0] = "Fahmi"
	arrNama[1] = "Jamaludin"
	arrNama[2] = "Mas'ud"

	fmt.Println(arrNama[0])
	fmt.Println(arrNama[1])
	fmt.Println(arrNama[2])

	fmt.Println("Cara 2")
	// deklarasi 2
	// var arrNama2 = [3]string{
	// 	"Nama1",
	// 	"Nama2",
	// 	"Nama3",
	// }
	arrNama2 := [4]string{
		"Nama1",
		"Nama2",
		"Nama3",
		"Nama4",
	}
	fmt.Println(arrNama2[0])
	fmt.Println(arrNama2[1])
	fmt.Println(arrNama2[2])

	fmt.Println(len(arrNama))
	fmt.Println(len(arrNama2))
}
