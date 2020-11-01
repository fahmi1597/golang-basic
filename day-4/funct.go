package main

import "fmt"

func iamf() {
	fmt.Println("ᓚᘏᗢ")
}

func setFullName(firstName string, lastName string) {
	fmt.Println("Full name :", firstName, lastName)
}

func getFullName() (string, string) {
	return "Fahmi", "Jamaludin"
}

func getFullName_() (firstName, lastName string) {
	firstName = "Fahmi"
	lastName = "Jamaludin"
	return
}

func main() {
	// function tanpa parameter
	iamf()
	// function + parameter
	lastName := "Jamaludin"
	setFullName("Fahmi", lastName)
	// function return value-multivalue
	firstName, lastName := getFullName()
	fmt.Println("Get Full name :", firstName, lastName)
	// function return value-multivalue
	a, b := getFullName_()
	fmt.Println(a, b)

}
