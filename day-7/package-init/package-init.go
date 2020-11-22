package main

import (
	"basic-golang/day-7/database"
	"fmt"
)

func main() {
	result := database.GetDatabase()
	fmt.Println(result)
}
