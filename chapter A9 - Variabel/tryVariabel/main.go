package main

import "fmt"

func main() {
	var firstName string = "john"
	lastName := "wick"

	fmt.Printf("Halo %s %s!\n", firstName, lastName)

	// ini metode deklarasi variabel dgn type inference
	one, two, three := 1, true, "three"
	fmt.Println(one, two, three)

	//variabel underscore utk menampung nilai tp gapingin dipake
	name, _ := "john", "wick"
	fmt.Println(name)

	//variabel new utk nampung variabel data bertipe pointer
	nama := new(string)
	fmt.Println(nama)

	var positiveNumber uint8 = 89
	var negativeNumber = -1243423644
	fmt.Printf("Bilangan positif: %d\n", positiveNumber)
	fmt.Printf("Bilangan negatif: %d\n", negativeNumber)

	fmt.Println("john wick")
fmt.Println("john", "wick")

fmt.Print("john wick\n")
fmt.Print("john ", "wick\n")
fmt.Print("john", " ", "wick\n")
	
	//
}