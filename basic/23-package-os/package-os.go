package main

import (
	"fmt"
	"os"
)

func main() {
	
	// args := os.Args
	// fmt.Println(args)
	if hostname, err := os.Hostname(); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println("Hostname : ", hostname)
		fmt.Println("Username : ", os.Getenv("USERNAME"))
	}

}
