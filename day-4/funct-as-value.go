package main

import "fmt"

func getGoodBye(name string) string {
	return "Good Bye, " + name
}

func main() {

	sayGB := getGoodBye
	result := sayGB("iamf")

	fmt.Println(sayGB("Fahmi"))
	fmt.Println(result)

}
