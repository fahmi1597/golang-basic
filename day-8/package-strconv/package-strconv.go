package main

import (
	"fmt"
	"strconv"
)

func main() {

	strToBool, err := strconv.ParseBool("true")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(strToBool)
	}

	strToInt, err := strconv.Atoi("52416483")
	if err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(strToInt)
	}

	intBase2ToStr := strconv.FormatInt(5, 2)
	fmt.Println(intBase2ToStr)
}
