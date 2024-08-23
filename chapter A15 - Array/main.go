package main

import "fmt"

func main() {
	var names [4]string
	names[0] = "trafalgar"
	names[1] = "tes 1"
	names[2] = "tes 2"
	names[3] = "tes 3"

	fmt.Println(names[0], names[1],names[2], names[3])

	//inisiasi nilai array awal
	var fruits = [4]string{"apel", "melon", "anggur", "semangka"}
	fmt.Println("Jumlah elemen \t\t", len(fruits))
	fmt.Println("Isi semua element \t", fruits)

	//inisiasi array tanpa jumlah elemen, jd arraynya bisa modular gitu lah
	var numbers = [...]int{2,3,2,4,3,4,5}

	fmt.Println("Data array: \t", numbers)
	fmt.Println("Jumlah elemen: \t", len(numbers))

	//array multidimensi
	var numbers1 = [2][3]int{[3]int{3,2,3}, [3]int{2,4,5}}
	var numbers2 = [2][3]int{{1,2,3},{7,4,6}}

	fmt.Println("numbers1", numbers1)
	fmt.Println("numbers2", numbers2)

	//Perulangan array menggunakan keyword for
	var buah = [4]string{"apple", "jeruk", "semangka", "durian"}

	for i:=0;i<len(buah);i++ {
		fmt.Printf("elemen %d : %s\n", i, buah[i])
	}

	//a15.7 Perulangan elemen array menggunakan keyword for-range
	var buah2 = [4]string{"melon", "nanas", "salak", "mangga"}
	for i, pruit := range buah2 {
		fmt.Printf("elemen %d : %s\n", i, pruit)
	}

	//a15.8 Penggunaan variabel underscore dlm for-range
	var fruitz = [3]string{"buah1","buah2","buah3"}
	for _,frut := range fruitz {
		fmt.Printf("Nama buah : %s\n", frut)
	}
	for i := range fruitz {
		fmt.Printf("Nomor buah : %d\n", i)
	}

	//A15.9 Alokasi elemen pake keyword MAKE
	var fruitss = make([]string, 2)
	fruitss[0] = "apple"
	fruitss[1] = "mango"

	fmt.Println(len(fruitss) , fruitss)
}