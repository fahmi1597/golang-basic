package main

import "fmt"

type Person interface {
	GetNama() string
	GetUmur() int
}

type Boy struct {
	Nama string
	Umur int
}

func (boy Boy) GetNama() string {
	return boy.Nama
}

func (boy Boy) GetUmur() int {
	return boy.Umur
}

func callPerson(person Person) {
	fmt.Println(person.GetNama())
	fmt.Println(person.GetUmur())
}

func main() {
	fahmi := Boy{
		Nama: "Fahmi",
		Umur: 23,
	}
	callPerson(fahmi)
}
