package main

import("fmt")

func main (){
	// string conversion
	var name string = "Fahmi"
	var s = name[0] // e --> byte --> uint8 --> 0~255
	var ss byte = 255

	fmt.Println(name)
	fmt.Println(s) 
	fmt.Println(string(s))
	//tanpa -1 => byte overflow 
	fmt.Println(string(ss)) 

}
