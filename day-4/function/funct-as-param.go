package main

import "fmt"

type wordFilter func(string) string

// func sayHelloWithFilter(name string, filter func(string) string) {
// 	filterName := filter(name)
// 	fmt.Println("Hello ", filterName)
// }
func sayHelloWithFilter(name string, wordFilter wordFilter) {
	filterName := wordFilter(name)
	fmt.Println("Hello ", filterName)
}

func doFilter(name string) string {
	if name == "anjing" {
		return "kasar"
	}
	return name
}

func main() {

	sayHelloWithFilter("anjing", doFilter)

}
