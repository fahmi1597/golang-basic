package main

import "fmt"

func main() {

	// penulisan 1 | mode while
	// counter := 1
	// for counter > 5 {
	// 	fmt.Println("Counter = ", counter)
	// 	counter++
	// }

	// penulisan 2 | dengan statement
	// for counter := 0; counter <= 10; counter++ {
	// 	fmt.Println("Counter = ", counter)
	// }

	// Mengakses array dengan for
	slice := []string{"A", "B", "C"}

	// cara 1
	// for i := 0; i < len(slice); i++ {
	// 	fmt.Println("Index ", i, ": ", slice[i])
	// }
	//cara 2 mirip for range | mirip foreach
	// for idx, value := range slice {
	// 	fmt.Println("Index ", idx, "= ", value)
	// }
	for _, value := range slice {
		fmt.Println(value)
	}

	person := map[string]string{
		"nama":    "Fahmi",
		"npm":     "52416483",
		"jurusan": "Teknik Informatika",
	}

	for key, value := range person {
		fmt.Println(key, ":", value)
	}

}
