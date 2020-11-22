package main

import "fmt"

func randomType() interface{} {
	return true
}

func main() {
	result := randomType()

	switch value := result.(type) {
	case string:
		fmt.Println("String", value)
	case int:
		fmt.Println("Int", value)
	default:
		fmt.Println("Unknown")
	}
}
