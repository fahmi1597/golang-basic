package main

import (
	"container/list"
	"fmt"
)

func main() {
	data := list.New()
	data.PushBack("1")
	data.PushBack("2")
	data.PushBack("3")
	data.PushFront("5")
	// fmt.Println(data.Front().Value)
	// fmt.Println(data.Back().Value)
	fmt.Println(data.Len())
	//iterasi dari depan
	for element := data.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value)
	}
}
