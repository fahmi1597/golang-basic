package main

import "fmt"

//NameServer is struct
type NameServer struct {
	IP string
}

func setNS(ns *NameServer) {
	ns.IP = "1.1.1.1"
}

func main() {

	// & get memory address
	// * read the value stored in the memory address
	nameServer1 := NameServer{
		IP: "8.8.8.8",
	}
	ptrNameServer1 := &nameServer1

	fmt.Println(*ptrNameServer1)
}

}
