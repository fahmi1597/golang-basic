package main

import "fmt"

// Person is orang
type Person struct {
	Name 	string
	Age  	int
}

func main() {

	var fahmi Person
	fahmi.Name = "Fahmi"
	fahmi.Age = 23
	
	// error jika tidak sesuai fieldnya.
	mikun := Person{"Mikun", 23}

	// Best practice
	budi := Person{
		Name: "Siapa Budi",
		Age:  20,
	}

	fmt.Println(fahmi)
	fmt.Println(budi)
	fmt.Println(mikun)
}
