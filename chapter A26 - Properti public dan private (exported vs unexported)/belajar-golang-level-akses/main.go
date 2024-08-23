package main

import (
	"belajar-golang-level-akses/library"
	"fmt"

	//A26.5 import dg prefix tanda titik
	. "belajar-golang-level-akses/library"
	//A26.6 Pemanfaatan alias saat import package
	f "fmt"
)

//A26
/*
	Property modifier public & private (penentu kpn suatu struct,
	fungsi, method bs diakses dr package lain dan kpn tdk)

	di go ga ada public & private sebenernya tp adanya:
	1. exported  -> public modifier -> komponen boleh diakses oleh package lain
	2. unexported -> private modifier -> komponen yg bs diakses dr file yg package nya sama. bs dr satu file yg sm atau
											file yg beda tp msh dlm satu package.


*/

//A26.3 Penggunaan package, import dan hak akses exported & unexported.

func main () {
	//cara ke 2 bakal bs krn modifier yg private itu dipanggil dulu ke function bermodifier publik
	library.SayHello("ethan")
	// library.introduce("ethan") //akan error krn akses modifiernya private

	var s1 = library.Student{"ethan", 21}
	fmt.Println("name", s1.Name)
	fmt.Println("grade", s1.Grade)

	/* 
		kesimpulan percobaan:
		utk menggunakan struct yg ada di package lain, selain nama struct nya hrs berbentuk exported
		properti yg diakses jg hrs exported jg.
	*/

	//A26.5 Import dg prefix tanda titik
	/*
		prefix tanda titik gunanya utk komponen yg berada di package lain yg diimport bs dijadikan se-level
		dg komponen package peng-import. Mksdnya selevel disini itu komponen import-an bs lgsg diakses tanpa perlu menuliskan nama package.
		seolah olah property tsb berada di file yg sama.

		tp teknik import ini bs jadiin kode jd ambigu. Makanya teknik import ini krg direkomen sih.
	*/

	var s2 = Student{"ishak", 23}
	fmt.Println("name:", s2.Name)
	fmt.Println("grade: ", s2.Grade)

	f.Printf("\n\n")
	f.Println("A26.6 Pemanfaatan alias saat import package")
	f.Println("Hellow world")

	//A26.7 Akses property dlm file yg package nya sama. 
	f.Printf("\n\n")
	f.Println("A26.7 Akses property dlm file yg package yg sm")
	sayHolla("ethan")

	//A26.8 Fungsi init()
	f.Printf("Name : %s\n", library.Student2.Name)
	f.Printf("Grade : %s\n", library.Student2.Grade)
}