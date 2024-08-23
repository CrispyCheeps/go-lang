package library

import "fmt"

//level aksesnya publik, ditandai dg nama fungsi diawali huruf besar.
func SayHello(name string) {
	fmt.Println("Hello")
	//cara ke-2
	introduce(name)
}

//level akses nya private krn nama fungsinya huruf kecil
func introduce(name string) {
	fmt.Println("nama saya", name)
}

type Student struct {
	//Klo properti nya diawali huruf kecil maka gabakal bs ke-export
	Name string
	Grade int
}

var Student2 = struct {
	Name string
	Grade int
} {}

func init() {
	Student2.Name = "John Wick"
	Student2.Grade = 2

	fmt.Println("--> library/library.go imported")
}