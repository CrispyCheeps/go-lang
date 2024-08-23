package main

import (
	"fmt"
	"math"
)

func calculate(d float64) (float64, float64) {
	//hitung luas
	var area = math.Pi * math.Pow(d/2,2)
	//hitung keliling
	var circumference = math.Pi * d

	//kembalikan 2 nilai
	return area, circumference
}

func main() {
	var diameter float64 = 15
	var area, circumference = calculate(diameter)

	fmt.Printf("luas lingkaran\t: %.2f \n", area)
	fmt.Printf("Keliling lingkaran\t: %.2f \n", circumference)
}

//A19.2 Fungsi dg predefined return value
func calculate2(d float64) (area float64, circumference float64) {
	area = math.Pi * math.Pow(d/2,2)
	circumference = math.Pi * d

	return
}

