package main

import "fmt"

func main() {
	var arrMonths = [...]string{
		"January",
		"February",
		"Maret",
		"April",
		"Mei",
		"Juni",
		"Juli",
		"Agustus",
		"September",
		"November",
		"Desember",
	}

	sliceArr1 := arrMonths[4:7]
	fmt.Print("Slice : ")
	fmt.Println(sliceArr1)
	fmt.Print("Length : ")
	fmt.Println(len(sliceArr1))
	fmt.Print("Capacity : ")
	fmt.Println(cap(sliceArr1))

}
