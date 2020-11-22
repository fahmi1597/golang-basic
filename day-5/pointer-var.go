package main

import "fmt"

type NameServer struct {
	IP string
}

func setNS(nameserver *string) {
	*nameserver = "1.1.1.1"
}

func main() {

	// & get memory address
	// * read the value stored in the memory address
	nameServer1 := NameServer{
		IP: "8.8.8.8",
	}
	nameServer2 := &nameServer1
	fmt.Println(nameServer1)
	fmt.Println(*nameServer2)
	// Paksa NS2 ke 1.1.1.1
	fmt.Printf("%T\n", &nameServer2.IP)
	setNS(&nameServer2.IP)
	fmt.Println(*nameServer2)

	nameServer3 := new(NameServer)
	fmt.Println(*nameServer3)
}

// 	fmt.Println("Name Server 1 Memory Address")
// 	fmt.Println(&nameServer1.IP)
// 	fmt.Println("Name Server 1 Value")
// 	fmt.Println(nameServer1.IP)
// 	fmt.Println("Name Server 2 Memory Address")
// 	fmt.Println(&nameServer2.IP)
// 	fmt.Println("Name Server 2 Value")
// 	fmt.Println(nameServer2.IP)
// }
