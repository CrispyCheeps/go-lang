package main

import "fmt"

//A24.1 Deklarasi Struct
type student struct {
	name string
	grade int
}

//A24.2 Penerapan Struct utk membuat object
func main() {
	var s1 student
	s1.name = "john wick"
	s1.grade = 2

	fmt.Println("name : ", s1.name)
	fmt.Println("grade : ", s1.grade)
	//A24.3 Inisialisasi Object Struct
	var siswa1 = student{}
	siswa1.name = "wick"
	siswa1.grade = 2

	var siswa2 = student{"ethan", 2}

	var siswa3 = student{name : "jason"}

	fmt.Println("student 1 :", siswa1.name)
	fmt.Println("student 2 :", siswa2.name)
	fmt.Println("student 3 :", siswa3.name)

	var s4 = student{name: "wayne", grade:2}
	var s5 = student{grade:2, name: "bruce"}
	fmt.Println("student 4 :", s4.name)
	fmt.Println("student 5 :", s5.name)

	//A24.4 Variabel Objek Pointer
	var ss1 = student{name : "wick", grade:2}
	var ss2 *student =&s1
	fmt.Println("student 1, name:", ss1.name)
	fmt.Println("student 4, name:", ss2.name)

	//A24.5 Embedded struct
	var z1 = studentt{}
	z1.name = "wick"
	z1.age = 21
	z1.grade = 2
	z1.person.age = 25

	var orang = person{name : "Ishak", age : 23}
	var mahasiswa = studentt{person: orang, grade: 8}

	fmt.Println("name : ", mahasiswa.name)
	fmt.Println("umur : ", mahasiswa.person.age)
	fmt.Println("grade : ", mahasiswa.grade)

	fmt.Println("name  :", z1.name)
    fmt.Println("age   :", z1.age)
    fmt.Println("age   :", z1.person.age)
    fmt.Println("grade :", z1.grade)


	//A24.8 Anonymous Struct
	/*
		anonymous struct -> struct yg tdk dideklarasikan di awal sbg tipe data baru, melainkan lgsg ketika pembuatan objek
		(+) cukup efisien klo digunakan pd use case pembuatan variabel objek yg struct-nya dipake sekali 
	*/

	fmt.Printf("\n\n")
	fmt.Println("A24.8 Anonymous Struct")

	/*
		salah satu aturan dlm pembuatan anonymous struct adalah,
		deklarasi hrs diikuti oleh inisialisasi. Nah tand {} di bawah adalah
		deklarasi struct tanpa pengisian property
	*/
	var s6 = struct{
		person
		grade int
	} {}

	// deklarasi struct dengan pengisian property
	var s7 = struct{
		person
		grade int
	} {
		person : person{"Ishak f", 23},
		grade : 9,
	}

	s6.person = person{"ishak", 23}
	s6.grade = 8

	fmt.Println("name :", s6.person.name)
	fmt.Println("name :", s6.person.age)
	fmt.Println("name :", s6.grade)
	fmt.Println("name :", s7.person.name)
	fmt.Println("name :", s7.person.age)
	fmt.Println("name :", s7.grade)

	//A24.9 Kombinasi slice & Struct
	fmt.Printf("\n\n")
	fmt.Println("A24.9 Kombinasi slice & Struct")
	var allStudents = []person{
		{name : "Ishak", age: 23},
		{name : "Tony", age : 24},
		{name : "Fajar", age : 23},
	}

	for _ , stu := range allStudents {
		fmt.Println(stu.name, "age is ", stu.age)
	}

	//A24.10 Inisialisasi Slice anonymous struct
	fmt.Printf("\n\n")
	fmt.Println("A24.10 Inisialisasi Slice anonymous struct")
	var allStudents2 = []struct {
		person
		grade int
	}{
		{person : person{"ishak", 23}, grade : 8},
		{person : person{"toney", 24}, grade : 3},
		{person : person{"fajar", 22}, grade : 4},
	}

	for _ , std := range allStudents2 {
		fmt.Println(std)
	}

	//A24.11 Deklarasi anonymous struct menggunakan keyword var
	//intinya ada cara lain selain penulisan struct menggunakan type
	// bedanya si s8 ini gbs digunakan kembali ga kyk type
	var s8 struct {
		person
		grade int 
	}

	s8.person = person{"ishak", 23}
	s8.grade = 8

	//A24.12 Nested struct
	type student struct {
		person struct {
			name string
			age int
		}
		grade int
		hobbies []string
	}

	//A24.15 Type alias
	fmt.Printf("\n\n")
	fmt.Println("A24.15 Type alias")
	fmt.Println(pp1)
	fmt.Println(pp2)



}

//A24.5 Embedded struct
type person struct {
	name string
	age int
}

type studentt struct {
	grade int
	person
	// A24.6 embedded struct dg nama property yg sama
	age int
}

//A24.13 Deklarasi dan inisialisasi struct secara horizontal
/*
	tanda semi colon (;) digunakan sbg pembatas deklarasi property yg dituliskan scr horizontal
 */
type person2 struct { name string; age int; hobbies []string}
var p1 = struct { name string; age int } { age: 22, name: "wick" }
var p2 = struct { name string; age int } { "ethan", 23 }

//A24.14 Tag property dalam struct
/*
	tag dimanfaatkan utk keperluan encode/decode data. info tag jg bs diakses lewat reflect.

 */
type person3 struct {
	name string `tag1`
	age int `tag2`
}

//A24.15 Type alias
// sebuah tipe data, spt struct. Bs dibuatkan alias baru. Caranya 
type Person struct {
	name string
	age int
}

type People = Person

var pp1 = Person{"wick", 21}
var pp2 = People{"tes", 23}








