package main

import "fmt"

func endAp() {

	if message := recover(); message != nil {
		fmt.Println("Message :", message)
	}
	fmt.Println("Aplikasi selesai")
}

func runApp(err bool) {
	defer endAp()
	if err {
		panic("Aplikasi error gatau knp")
	}
	fmt.Println("Aplikasi berjalan")
}

func main() {
	runApp(true)
	runApp(false)
}
