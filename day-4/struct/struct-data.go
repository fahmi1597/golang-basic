package main

import "fmt"

// Person is orang
type Person struct {
	Name string
	Age  int
}

func main() {

	var fahmi Person
	fahmi.Name = "Fahmi Jamaludin"
	fahmi.Age = 23

	budi := Person{
		Name: "Budi",
		Age:  20,
	}
	//error jika tidak sesuai fieldnya.
	mikun := Person{"Mikun", 23}

	fmt.Println(fahmi)
	fmt.Println(budi)
	fmt.Println(mikun)
}
