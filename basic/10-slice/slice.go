package slice

import "fmt"

func main() {

	// both of these are array
	arr1 := [3]int{1, 2, 3}
	var arr2 = [...]int{1, 2, 3}
	fmt.Println(arr1)
	fmt.Println(arr2)

	// slice
	// slice[start:end]
	// slice[start:]
	// slice[:end]
	// slice[:]
	slice := []int{1, 2, 3}
	fmt.Println(slice)

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

	fmt.Println(arrMonths)
	sliceOfMonths := arrMonths[4:7]
	fmt.Print("Slice : ")
	fmt.Println(sliceOfMonths)
	fmt.Print("Length : ")
	fmt.Println(len(sliceOfMonths))
	fmt.Print("Capacity : ")
	fmt.Println(cap(sliceOfMonths))
	fmt.Println()
	sliceOfSlice := append(sliceOfMonths, "Added")
	fmt.Print("New Slice : ")
	fmt.Println(sliceOfSlice)
	fmt.Print("Length : ")
	fmt.Println(len(sliceOfSlice))
	fmt.Print("Capacity : ")
	fmt.Println(cap(sliceOfSlice))
	fmt.Println()
	fmt.Println(sliceOfMonths)
	sliceOfSlice[3] = "Replaced"
	fmt.Println("State Now")
	fmt.Println(sliceOfMonths)
	fmt.Println(arrMonths)
}
