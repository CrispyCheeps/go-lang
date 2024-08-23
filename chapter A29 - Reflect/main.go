package main

import (
	"fmt"
	"reflect"
)

//A.29.2. Informasi property variabel objek
type student struct {
	Name string
	Grade int
}

//method buat struct student (nampilin info tiap property struct student)
func (s *student) getPropertyInfo() {
	var reflectValue = reflect.ValueOf(s)

	//dilakukan pengecekan apakah reflectValue pointer apa bkn
	//jika iya maka perlu diambil objek aslinya pake .Elem()
	if reflectValue.Kind() == reflect.Ptr {
		reflectValue = reflectValue.Elem()
	}

	var reflectType = reflectValue.Type()

	//.Field ini mengembalikan info tiap property struct diambil berurutan.
	for i := 0; i < reflectValue.NumField(); i++ {
		fmt.Println("nama : ", reflectType.Field(i).Name)
		fmt.Println("tipe data : ", reflectType.Field(i).Type)
		fmt.Println("nilai : ", reflectValue.Field(i).Interface())
		fmt.Println(" ")
	}
}

//A29.3 Pengeksesan info method variabel objek
func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("tipe variabel : ", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("Nilai variabel :", reflectValue.Interface())
	}

	var nilai = reflectValue.Interface().(int)
	fmt.Println(nilai)

	//Uji coba suksesi method A29.2
	var s1 = &student{Name: "John wick", Grade: 2}
	s1.getPropertyInfo()

	//Uji coba suksesi method A29.3
	fmt.Println("A29.3")
	fmt.Println("nama: ", s1.Name)

	var reflectValue2 = reflect.ValueOf(s1)
	var method = reflectValue2.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("wick"),
	})

	fmt.Println("nama : ", s1.Name)
}
