package main

import "fmt"

func main() {
	//A16.1 Inisialisasi slice
	/*
		Perbedaan slice dan array yg bs diketahui saat pendeklarasian variabel adalah
		jika jumlah elemen tdk dituliskan, maka variabel tersebut adalah slice.
		array = kumpulan nilai / elemen
		slice = referensi dr kumpulan nilai tsb
	*/
	var fruits = []string{"apple", "grape", "banana", "melon"}
	fmt.Println(fruits[0])

	var _ = []string{"apple", "grape"} //slice
	var _ = [2]string{"banana", "melon"} //array

	//A16.2 Hubungan slice dgn array & operasi slice
	var newFruits = fruits[2:]
	fmt.Println(newFruits)

	//A16.3 Slice merupakan tipe data reference
	var aFruits = fruits[0:3]
	var bFruits = fruits[0:4]

	var aaFruits = aFruits[0:1]
	var bbFruits = bFruits[0:2]

	fmt.Println("A16.3")
	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(bbFruits)

	bbFruits[0] = "pinnaple"
	fmt.Println("A16.3")
	fmt.Println(fruits)
	fmt.Println(aFruits)
	fmt.Println(bFruits)
	fmt.Println(aaFruits)
	fmt.Println(bbFruits)
	/*
		Semua elemen yg awalnya apple akan berubah jd pinnaple krn mempunya referensi yg sama
		(sifat dasar dari slice)
	*/

	//A16.4 Fungsi len() -> buat hitung jumlah elemen di slice atau array
	var buah = []string{"melon", "mangga", "apel", "salak"}
	fmt.Println("A16.4")
	fmt.Println(len(buah))

	//A16.5 Fungsi cap() -> utk ngetahuin lebar atau kapasitas maksimum dr slice / array
	fmt.Println("")
	fmt.Println("A16.5")
	fmt.Println(len(buah))
	fmt.Println(cap(buah))

	var aBuah = buah[0:3]
	fmt.Println(len(aBuah))
	fmt.Println(cap(aBuah))

	var bBuah = buah[1:2]
	fmt.Println(len(bBuah))
	fmt.Println(cap(bBuah))

	//A16.6 Fungsi Append
	var cBuah = append(buah, "pepaya")
	fmt.Println(buah)
	fmt.Println(cBuah)
	/*
		si Append ini ada punya kondisi:
		1. Kalau len == cap , maka elemen append() nya bakal jd referensi baru
		2. Kalau len < cap , maka elemen append() nya bakal msk ke referensi yg sm (yg kosong)
	*/
	fmt.Println("")
	fmt.Println("praktikum append")

	var bbuah = buah[0:2]

	fmt.Println(cap(bbuah))
	fmt.Println(len(bbuah))
	fmt.Println(buah)
	fmt.Println(bbuah)

	var cbuah = append(bbuah, "pepayaa")

	fmt.Println(buah)
	fmt.Println(bbuah)
	fmt.Println(cbuah)
	fmt.Printf("\n\n")

	//A16.7 Fungsi Copy -> utk mengcopy parameter ke 2 (src) ke parameter pertama (dst)
	//copy(dst, src)
	fmt.Println("A16.7 Copy")
	dst := make([]string, 3)
	src := []string{"melon", "nanas", "apel", "jeruk"}
	n := copy(dst, src)

	fmt.Println(dst)
	fmt.Println(src)
	fmt.Println(n)
	//berikut contoh slice yg udh ada isinya:
	fmt.Println("slice yg udh ada isinya")
	dst2 := []string{"kentang", "kentang", "kentang"}
	src2 := []string{"melon", "nanas", "apel", "jeruk"}
	n2 := copy(dst2, src2)

	fmt.Println(dst2)
	fmt.Println(src)
	fmt.Println(n2)

	//A16.8 Pengaksesan Elemen slice dengan 3 indeks
	//contoh fruits[0:1:1] -> indeks ke 3 itu nunjukin kapasitas slice yang akan dislicing.

	fmt.Print("\n")
	fmt.Println("A16.8")
	var A168Buah = []string{"apple", "grape", "banana"}
	var aA168Buah = A168Buah[0:2]
	var bA168Buah = A168Buah[0:2:2]

	fmt.Println(A168Buah)
	fmt.Println(len(A168Buah))
	fmt.Println(cap(A168Buah))

	fmt.Println(aA168Buah)
	fmt.Println(len(aA168Buah))
	fmt.Println(cap(aA168Buah))

	fmt.Println(bA168Buah)
	fmt.Println(len(bA168Buah))
	fmt.Println(cap(bA168Buah))

}