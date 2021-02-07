package main

import (
	"fmt"
	"math"
)

func main() {

	var diameter float64 = 15
	var luas, keliling = calculate(diameter)

	fmt.Printf("luas lingkaran\t\t: %.2f \n", luas)
	fmt.Printf("keliling lingkaran\t: %.2f \n", keliling)
}

func calculate(d float64) (luas float64, keliling float64) {
	luas = math.Pi * math.Pow(d/2, 2)
	keliling = math.Pi * d

	return
}
