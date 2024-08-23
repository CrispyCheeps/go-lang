package main

import "fmt"

func main() {
	const firstName string = "john"
	fmt.Print("halo ", firstName, "!\n")

	//deklarasi multi konstanta
	const(
		square = "kotak"
		isToday bool = true
		numeric uint8 = 1
		floatNum = 2.2
	)

	//contoh deklrasi utk nilai yg sama
	const (
		a = "konstanta"
		b
	)

	//deklarasi multi const dlm satu baris (type manifest typing method
	const satu, dua = 1, 2

	//deklarasi multi const dlm satu baris (type inference method)

	

}