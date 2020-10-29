
package main

import ("fmt")

func main(){

	// # Cara deklarasi pertama
	var name string

	name = "Fahmi"
	fmt.Println(name)
	
	//  Cara deklarasi + inisialisasi
	var name_2 = "みくん"
	fmt.Println(name_2)

	// # Cara deklarasi tanpa var
	name_3 := "IamF"
	fmt.Println(name_3)
		
		// # Overwrite/modification
		name_3 = "IamF - Overwritted"
		fmt.Println(name_3)
}