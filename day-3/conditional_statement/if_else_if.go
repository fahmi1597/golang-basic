package main

import "fmt"

func main() {
	username := "Fahmi"

	if username == "Fahmi" {
		fmt.Println("Logged in as", username)
	} else if username == "" {
		fmt.Println("Username not found")
	} else {
		fmt.Println("Access Denied")
	}

	if pass := int8(len(username)); pass > 7 {
		fmt.Println("Strong Password")
	} else {
		fmt.Println("Weak Password")
	}

}
