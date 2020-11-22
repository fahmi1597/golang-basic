package main

import (
	"flag"
	"fmt"
)

func main() {
	//merupakan pointer
	hostname := flag.String("hostname", "example.com", "Put your host")
	username := flag.String("username", "bob", "Put your username")
	password := flag.String("password", "alice", "Put your password")

	flag.Parse()

	fmt.Println("Hostname \t :", *hostname)
	fmt.Println("Username \t :", *username)
	fmt.Println("Password \t :", *password)
}
