package main

import (
	"fmt"
	"runtime"
)

//A30 Goroutine

/*
	Goroutine =mirip= thread?

	Concurrent programming vs paralel
	konkurensi -> komposisi dari sebuah proses
	paralel -> bagaimana eksekusinya berlangsung
*/

func print(till int, message string) {
	for i := 0; i < till; i++ {
		fmt.Println((i+1), message)
	}
}

func main() {

	//digunakan utk menentukan jumlah core yg diaktifkan utk eksekusi program
	runtime.GOMAXPROCS(2)
	
	//penerapan gourutin hrs nambahin go di dpn pemanggilan fungsi kek gini
	go print(5, "halo")
	print(5, "apa kabar")

	var input string
	fmt.Scanln(&input)

	//Uji coba fmt.Scanln
	var s1, s2, s3 string
	fmt.Scanln(&s1, &s2, &s3)

	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3)
}