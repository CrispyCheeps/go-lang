package main

import "fmt"

type student struct {
	name string
	height float64
	age int32
	isGraduated bool
	hobbies []string
}

var data=student{
	name:        "wick",
    height:      182.5,
    age:         26,
    isGraduated: false,
    hobbies:     []string{"eating", "sleeping"},
}

func main () {
	//%b buat ngubah numerik jd string berbentuk biner.
	fmt.Printf("%b\n", data.age)

	//ngeformat data numerik yg merupakan kode unicode, jd string karakter unicodenya.
	fmt.Printf("%c\n", 1235)

	//%d ngeformat data numerik mjd bilangan berbasis 10 (bilangan yg kita gunakan sekarang)
	fmt.Printf("%d\n", data.age)

	//%e/ %E utk ngubah data numerik desimal mjd numerik standar scientific notation.
	//1.825000E+02 maksudnya adalah 1.825 x 10^2, dan hasil operasi tersebut adalah sesuai dengan data asli = 182.5.
	fmt.Printf("%E\n", data.height)

	//%F format numerik desimal, dg lebar bisa ditentukan. Lebar defaultnya 6 dibelakang koma
	fmt.Printf("%f\n", data.height)
	fmt.Printf("%.2f\n", data.height)

	fmt.Println("%G")
	//%G mirip %F tp persen G biasa dipake utk data yg jumlah digit desimalnya cukup banyak. Dan %G ini gabisa dicustom kyk %F
	fmt.Printf("%e\n", 0.123123123123)
	// 1.231231e-01
	fmt.Printf("%f\n", 0.123123123123)
	// 0.123123
	fmt.Printf("%g\n", 0.123123123123)


	//%P digunakan utk nampilin alamat pointer. Diturliskan dlm btk numerik berbasis 16 dgn prefix 0x
	fmt.Printf("%p\n", &data.name)

	//%q digunakan utk escape string
	fmt.Printf("%q\n", `" name \`)

	//%s utk format data string
	fmt.Printf("%s\n", data.name)

	//%t digunakan utk format data boolean, nampilin nilai bool
	fmt.Printf("%t\n", data.isGraduated)

	//%T utk ngambil tipe variabel yg akan diformat.
	fmt.Printf("%T\n", data.name)
	// string
	fmt.Printf("%T\n", data.height)
	// float64
	fmt.Printf("%T\n", data.age)
	// int32
	fmt.Printf("%T\n", data.isGraduated)
	// bool
	fmt.Printf("%T\n", data.hobbies)
	// []string

	fmt.Println("%v")
	//%v utk memformat data apa aja termasuk interface{} dg pengembalian string nilai data aslinya.
	fmt.Printf("%v\n", data)

	fmt.Println("%+v")
	//%+v utk memformat struct return nya nama tiap property dan nilainya scr urut dg struktur struct.
	fmt.Printf("%+v\n", data)

	//%#v return format struct (nilai tiap property) + bgmn objek tsb dideklarasikan
	fmt.Println("%#v")
	fmt.Printf("%#v", data)


}

