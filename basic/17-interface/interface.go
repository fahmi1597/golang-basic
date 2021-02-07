package main

import "fmt"

type Animal interface {
	GetNama() int
	GetRas() string
}

type Cat struct {
	Nama string
	Ras  string
}

func (c *Cat) GetNama() string {
	return c.Nama
}

func (c *Cat) GetRas() string {
	return c.Ras
}

func describe(c Cat) {
	fmt.Println("Nama:", c.GetNama())
	fmt.Println("Ras", c.GetRas())
}

func main() {
	meng := Cat{
		Nama: "Meng",
		Ras:  "OrenBelang",
	}
	describe(meng)
}
