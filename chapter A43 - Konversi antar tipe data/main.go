package main

import (
	"fmt"
	"strconv"
)

func main() {
	var str = "124"
	fmt.Printf("%T", str)
	var num, err = strconv.Atoi(str)

	if err == nil {
		fmt.Println(num)
		fmt.Printf("%T", num)
	}

	//A.43.2. Konversi Data Menggunakan Teknik Casting
	// konversi nilai 24 bertipe int ke float64
	var a float64 = float64(24)
	fmt.Println(a) // 24

	// konversi nilai 24.00 bertipe float32 ke int32
	var b int32 = int32(24.00)
	fmt.Println(b) // 24

	//A.43.3. Casting string ↔ byte
	fmt.Println("A.43.3. Casting string ↔ byte")
	text1 := "Halo"
	b_ := []byte(text1)
	fmt.Println(b_)

	//A.43.4. Type Assertions Pada Tipe any atau Interface Kosong (interface{})

}