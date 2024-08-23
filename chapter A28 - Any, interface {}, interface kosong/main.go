//A28 Any / interface {} / interface kosong
/*
	interface kosong mrpkn tipe data yg sgt spesial krn variabel ini bs
	menampung sgala jenis data.
*/

// A28.1 Penggunaan any / interface {}
package main

import (
	"fmt"
	"strings"
)

//A28.4 Casting Variabel interface kosong ke objek pointer
type person struct {
	name string
	age int
}

func main () {
	var secret interface{}

	secret = "ethan hunt"
	fmt.Println(secret)

	secret = []string{"apple", "manggo", "banana"}
	fmt.Println(secret)

	secret = 12.4
	fmt.Println(secret)

	//A28.3 casting variabel any / interface kosong
	// teknik casting pada any ini dinamai -> type assertions
	var secret2 interface{}

	secret2 = 2
	var number = secret2.(int) * 10
	fmt.Println(secret, "multiplied by 10 is: ", number)

	secret2 = []string{"apple", "manggo", "banana"}
	var gruits = strings.Join(secret2.([]string), ", ")
	fmt.Println(gruits, "is my favourite fruits")

	//A28.4 Casting Variabel interface kosong ke objek pointer
	var secret3 interface{} = &person{name: "wick", age:27}
	var name = secret3.(*person).name
	fmt.Println(name)

	//A28.5 Kombinasi slice, map dan interface{}
	//[]map[string]interface{} -> sering dipake utk nyimpen sekumpulan data berbasis key-value

	var person = []map[string]interface{} {
		{"name": "wick", "age":23},
		{"name" : "Ethan", "age": 23},
		{"name" : "Bourne", "age": 22},
	}

	for _, each := range person {
		fmt.Println(each["name"], "age is", each["age"])
	}

	var fruits = []interface{}{
		map[string]interface{}{
			"name": "strawberry",
			"total": 10},
			[]string{"manggo", "pineapple", "papaya"},
			"orange", 
		}
	
	for _, each := range fruits {
		fmt.Println(each) } 
}
