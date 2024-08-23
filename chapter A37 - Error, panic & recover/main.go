package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

//A37.4 - Penggunaan Recover
/*
	recover berguna utk menghandle panic error
	jd, ketika muncul panic dan ada recover, program jd ga berhenti.
	panic akan mereturn nill kalau panic ga ada.
*/
func catch() {
	if r:=recover() ; r != nil {
		fmt.Println("error occured", r)
	} else {
		fmt.Println("Application running perfectly!")
	}
}

func main() {

	//A37.4 - Penggunaan Recover
	defer catch()

	var input string
	fmt.Print("Type some number: ")
	fmt.Scanln(&input)

	var number int
	var err error

	//strconv.Atoi() konversi jd numerik
	// return 1 -> hasil konversi. Return 2 -> error
	number, err = strconv.Atoi(input)

	if err == nil {
		fmt.Println(number, "is number")
	} else {
		fmt.Println(input, "is not number")
		fmt.Println(err.Error())
	}
	//A37.2 Membuat custom error
	var name string
	fmt.Print("type your name: ")
	fmt.Scanln(&name)

	if valid, err := validate(name); valid{
		fmt.Println("halo", name)
	} else {
		//A37.3 Penggunaan Panic 
		panic(err.Error())
		fmt.Println("end")
	}
}

//A37.2 Membuat custom error
func validate(input string) (bool, error) {
	if strings.TrimSpace(input) == "" {
		return false, errors.New("cannot be empty")
	}
	return true, nil
}
