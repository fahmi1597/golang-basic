package main

import "fmt"

func iamf() {
	fmt.Println("ᓚᘏᗢ")
}

func setFullName(firstName string, lastName string) {
	fmt.Println("Full name :", firstName, lastName)
}

func getFullName() (string, string) {
	return "Fahmi", "J"
}

func getFullName_() (firstName, lastName string) {
	firstName = "Fahmi"
	lastName = "J"
	return
}

func main() {
	// function tanpa parameter
	iamf()

	// function + parameter (fullname, lastname)
	lastName := "J"
	setFullName("Fahmi", lastName)
	
	// function multivalue return 
	firstName, lastName := getFullName()
	fmt.Println("Get full name :", firstName, lastName)
	
	// function multivalue return 
	a, b := getFullName_()
	fmt.Println(a, b)

}
