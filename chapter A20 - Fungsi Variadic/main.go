package main

import (
	"fmt"
	"strings"
)

//A20.1 Penerapan fungsi variadic -> fungsi dg parameter bs menampung nilai sejenis yg tdk terbatas jumlahnya

//pokoknya perlakuan fungsinya ya kek fungsi biasa, tp jumlah parameter yg dimasukin bs banyak
func main() {
	var avg = calculate(2,4,3,5,4,3,3,5,5,3)
	var msg = fmt.Sprintf("Rata-rata : %.2f" ,avg) // Sprintfmengembalikan nilai dlm btk string
	fmt.Println(msg)

	//A20.3 Pengisian parameter fungsi variadic menggunakan data slice
	var numbers = []int{2,4,3,5,4,3,3,5,5,6}
	var avg2 = calculate(numbers...)
	var msg2 = fmt.Sprintf("Rata-rata : %.2f", avg2)

	fmt.Println(msg2)

	/*
		var numbers = []int{2, 4, 3, 5, 4, 3, 3, 5, 5, 3}
		var avg = calculate(numbers...)

		atau

		var avg = calculate(2, 4, 3, 5, 4, 3, 3, 5, 5, 3)
	*/

	yourHobbies("wick", "sleeping", "eating")
	//atau
	var hobbies = []string{"sleeping", "eating"}
	yourHobbies("wicky", hobbies...)

}

func calculate(numbers ...int) float64 {
	var total int = 0
	for _,number := range numbers {
		total += number
	}

	var avg = float64(total) / float64(len(numbers))
	return avg
}

//A20.4 fungsi dgn parameter biasa & variadic
func yourHobbies(name string, hobbies ...string) {
	var hobbiesAsString = strings.Join(hobbies, ", ")

	fmt.Printf("Hellow, my name is : %s\n", name)
	fmt.Printf("My hobbies are: %s\n", hobbiesAsString)
}

