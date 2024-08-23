package main

import (
	"fmt"
	"strings"
)

type student struct {
	name string
	grade int
}

//A25.1 Penerapan method
func (s student) sayHello() {
	fmt.Println("Hellow", s.name)
}

//receiver / nama fungsi / parameter / returnType
// receriver ini menunjukan bahwa method getNameAt adalah method dr type student
func (s student) getNameAt (i int) string {
	return strings.Split(s.name, " ")[i-1]
}


//A25.2 Method Pointer 
func(s student) changeName1(name string) {
	fmt.Println("----> on changeName 1, name changed to", name)
	s.name = name
}

//Method pointer
func(s *student) changeName2(name string) {
	fmt.Println("-----> on changeName2, name changed to", name)
	s.name = name
}


func main() {
	var s1 = student{"ishak Febrianto", 23}
	s1.sayHello()

	var name = s1.getNameAt(1)
	fmt.Println("nama panggilan : ", name)

	//A25.2 Method Pointer 
	/*
		Kesimpulannya kalo mau bisa switch2 value pake method pointer ini aja
	 */
	fmt.Printf("\n\n")
	fmt.Println("A25.2 Method Pointer ")
	var s2 = student{"john wick", 21}
	fmt.Println("s2 before", s2.name)

	s2.changeName1("jason bourne")
	fmt.Println("s1 after changedName1", s2.name)

	s2.changeName2("ethan hunt")
	fmt.Println("s1 after changedName2", s2.name)

	/*
		Keistimewaan method pointer, method itu bs dipanggil dr objek pointer
		maupun objek biasa.
	 */
	 //pengaksesan method dr variabel objek biasa
	 var s3 = student{"joko", 42}
	 s3.sayHello()
	 fmt.Println(s3)
	 fmt.Println(&s3)
	 fmt.Printf("%p\n",s3)

	 var s4 = &student{"etes", 5}
	 s4.sayHello()
	 fmt.Printf("%p\n",s4)
	 fmt.Println(*s4)
}





