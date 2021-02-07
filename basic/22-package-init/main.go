package main

import (
	"fmt"
	"golang-basic/day-7/package-init/database"
	)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
