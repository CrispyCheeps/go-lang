package main

import "fmt"

func main() {
	//string
	var xs = "123"
	for i, v := range xs{
		fmt.Println("Index=", i, "Value=", v)
	}

	//array
	var ys = [5]int{10,20,30,40,50}
	for _, v := range ys {
		fmt.Println("Value=", v)
	}

	//slice
	var zs = ys[0:2]
	for _,v := range zs {
		fmt.Println("Value slice=", v)
	}


	//map
	var kvs = map[byte]int{'a':0, 'b':1, 'c':2}
	for k, v := range kvs {
		fmt.Println("Key=", k, "Value=", v)
	}

	// boleh juga baik k dan atau v nya diabaikan
	for range kvs {
    	fmt.Println("Done")
	}

	// selain itu, bisa juga dengan cukup menentukan nilai numerik perulangan
	for i := range 5 {
    	fmt.Print(i) // 01234	
	}

	fmt.Printf("\n")

	for i := 1; i<=10; i++ {
		if i % 2 ==1 {
			continue
		}

		if i>8 {
			break
		}

		fmt.Println("Angka", i)
	}

	//pemanfaatan label dlm perulangan bersarang
	outerLoop:
	for i := 1; i < 5; i++ {
		for j :=0; j<5; j++ {
			if i==3 {
				break outerLoop
			}
			fmt.Print("matriks [",i, "] [",j, "]","\n")
		}
	}
}