package main

import "fmt"

func main() {

	person := map[string]string{
		"nama":    "Fahmi",
		"npm":     "52416483",
		"jurusan": "Teknik Informatika",
	}

	fmt.Println(person)
	person["gelar"] = "S1"
	fmt.Println(person)
	delete(person, "gelar")
	fmt.Println(person)

	book := make(map[int]string)
	book[0] = "Cyber Security"
	book[1] = "Bash Scripting"

	fmt.Println(book)
	fmt.Println(book[1])

}
