package main

import (
	"fmt"
)

// error
// BASE_PATH := "/root"

// Ok! unused
var BASE_PATH string = "/root"


func main() {

	// Topics : Deklarasi Variable
	
	// 1. Deklarasi disertakan dengan tipe data, lalu assign nilai.
	
	var namaDepan string
	namaDepan = "Fahmi"
	fmt.Println(namaDepan)
	
	// 2. Deklarasi dan inisialisasi nilai, tipe data ditentukan dari nilai yang assign

	var namaPanggilan = "みくん"
	fmt.Println(namaPanggilan)

	// 3. Deklarasi pendek (short assignment), *hanya bisa dilakukan didalam fungsi*

	namaInGame := "IamF"
	fmt.Println(namaInGame)

	// 4. Multi deklarasi variable

	var (
		firstName = "Siapa"
		lastName = "Saya"
	)
	fmt.Println(firstName)
	fmt.Println(lastName)

}