package main

import (
	"fmt"
	"os"
)

/*
	buat nunda eksekusi line code jg akhir2 aja
*/
func orderSomeFood(menu string) {
	defer fmt.Println("Terimakasih, silahkan tunggu!")
	if menu == "pizza" {
        fmt.Print("Pilihan tepat!", " ")
        fmt.Print("Pizza ditempat kami paling enak!", "\n")
        return
    }
	fmt.Println("Pesanan Anda:", menu)
}

func main() {
	// defer fmt.Println("halo")
	// fmt.Println("Selamat datang")

	// orderSomeFood("pizza")

	//A36.2 Kombinasi defer dan IIFE
	/*
		eksekusi defer itu di akhir blok fungsi bkn blok lainnya, misal seleksi kondisi.
		Nah supaya bs dieksekusi 
	*/
	number := 3
	if number == 3 {
		fmt.Println("halo 1")
		//skrg defer bisa di dlm blok if else
		defer fmt.Println("halo 3")
	}

	fmt.Println("halo 2")

	//A36.3 Penerapan fungsi os.Exit()
	defer fmt.Println("halo")
	os.Exit(1)
	fmt.Println("selamat datang")
}