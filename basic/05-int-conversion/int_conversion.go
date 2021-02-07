package main

import (
	"fmt"
)

func main() {

	var num32 int32 = 255
	var num64 int64 = int64(num32)
	var num8 int8 = int8(num32)
	var unum8 uint8 = uint8(num32)

	fmt.Println(num32)
	fmt.Println(num64)

	//overflow
	fmt.Println(num8)
	fmt.Println(unum8)

}
