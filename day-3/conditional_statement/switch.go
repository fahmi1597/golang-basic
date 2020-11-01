package main

import "fmt"

func main() {
	name := "iamf"

	switch name {
	case "fahmi":
		fmt.Println("Bukan Fahmi")
	case "iamf":
		fmt.Println("iamf")
	default:
		fmt.Println("Males")
	}

	//short switch
	switch newName := name[1:3]; len(newName) < len(name) {
	case true:
		fmt.Println("Ok", newName)
	case false:
		fmt.Println("Not Ok,", newName)
	}

	// no expression switch
	switch {
	case name == "mikun":
		fmt.Println("Wrong person")
	case name != "mikun":
		fmt.Println("Right person")
	default:
		fmt.Println("Wrong person again")

	}

}
