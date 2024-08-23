package main

import "fmt"

/*
	closure -> anonym func yg disimpan dlm variabel (fungsi tanpa nama lah intinya)
	dgn closure kita bs mendesain :
		1. Fungsi dlm fungsi
		2.membuat fungsi yg ngembaliin fungsi
	biasa dimanfaatkan utk ngebungkus suatu proses yg hanya dijalankan sekali aja
	atau hny dipakai pd blok tertentu aja.
*/

func main() {
	var getMinMax = func(n []int) (int, int) {
		var min, max int
		for i, e := range n {
			switch {
			case i == 0:
				max, min = e, e
			case e > max:
				max = e
			case e < min :
				min = e
			}
		}

		return min, max
	}

	var numbers = []int{2,3,4,3,4,2,3}
	var min, max = getMinMax(numbers)
	fmt.Printf("data : %v\nmin : %v\nmax : %v\n", numbers,min,max)

	//Implementasi dari A21.3
	var max2 = 3
	var numbers2 = []int{2,3,0,4,3,2,0,4,2,0,3}
	var howMany, getNumbers = findMax(numbers2, max2)
	var theNumbers = getNumbers()

	fmt.Println("numbers\t:", numbers2)
	fmt.Printf("find \t: %d\n\n", max2)

	fmt.Println("found \t:", howMany)
	fmt.Println("value \t:", theNumbers)

}

//A21.3 Closure sbg nilai kembalian

/*
	closure bs dijadikan sbg nilai kembalian sbuah fungsi
*/

func findMax(numbers []int, max int) (int, func() []int) {
	var res []int
	for _,e := range numbers {
		if e <= max {
			res = append(res, e)
		}
	}
	return len(res), func() []int {
		return res
	}
}