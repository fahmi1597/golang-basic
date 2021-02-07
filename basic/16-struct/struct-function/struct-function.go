package main

import "fmt"

// Mobil is not orang
type Mobil struct {
	Merk  string
	Tahun string
}

func (mobil Mobil) getInfo() {
	fmt.Println("Merk", mobil.Merk)
	fmt.Println("Tahun", mobil.Tahun)
}
func (mobil Mobil) maju() {
	fmt.Println("Mobil maju")
}
func (mobil Mobil) operator(name string) {
	fmt.Println("Mobil", mobil.Merk, "oleh", name)
}

func main() {

	mobil1 := Mobil{
		Merk:  "BMW",
		Tahun: "2000",
	}

	// getInfo(Mobil1)
	// maju(Mobil1)
	mobil1.operator("fahmi")
	mobil1.getInfo()
	mobil1.maju()

}
