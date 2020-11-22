package main

import "fmt"

//NameServer is struct
type NameServer struct {
	IP string
}

func (ns *NameServer) setNS() {
	ns.IP = "1.1.1.1"

}
func main() {
	// & get memory address
	// * read the value stored in memory
	nameServer1 := NameServer{
		IP: "8.8.8.8",
	}
	nameServer1.setNS()
	fmt.Println(nameServer1)
}
