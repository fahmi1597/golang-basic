package main

import (
	"fmt"
	"regexp"
)

func main() {
	regex := regexp.MustCompile("[a-z]ahmi")

	fmt.Println(regex.MatchString("fahmi"))
	fmt.Println(regex.MatchString("zahmi"))
	fmt.Println(regex.MatchString("Fahmi"))
}
