package main

import (
	"encoding/json"
	"fmt"
)

//A53.1 Decode JSON ke variabel Objek Struct

type User struct {
	FullName string `json:"Name"`
	Age int
}

func main() {
	var jsonString = `{"Name": "john wick", "Age" : 27}`
	var jsonData = []byte(jsonString)

	var data User

	/*
	Fungsi unmarshal hanya menerima data json dalam bentuk []byte ,
	maka dari itu data json string perlu di-casting terlebih dahulu ke tipe ,
	sebelum akhirnya digunakan pada pemanggilan fungsi .
	*/
	var err = json.Unmarshal(jsonData, &data)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user :", data.FullName)
	fmt.Println("age :", data.Age)


	//A.53.2. Decode JSON Ke map[string]interface{} & interface{}
	fmt.Println("//A.53.2. Decode JSON Ke map[string]interface{} & interface{}")
	var data1 map[string]interface{}
	json.Unmarshal(jsonData, &data1)

	fmt.Println("user :", data1["Name"])
	fmt.Println("age :", data1["Age"])

	var data2 interface{}
	json.Unmarshal(jsonData, &data2)

	var decodeData = data2.(map[string]interface{})
	fmt.Println("user :", decodeData["Name"])
	fmt.Println("age :", decodeData["Age"])

	//A.53.3. Decode Array JSON Ke Array Objek
	fmt.Println("A.53.3. Decode Array JSON Ke Array Objek")
	var jsonString2 = `[
		{"Name": "ishak obriant", "Age": 27},
		{"Name": "ethan hunt", "Age": 32}
	]`

	var data3 []User

	var err2 = json.Unmarshal([]byte(jsonString2), &data3)
	if err2 != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Println("user 1: ", data3[0].FullName)
	fmt.Println("user 2: ", data3[1].FullName)

	//A.53.4 Enocde Objek ke JSON String
	var object = []User{{"john wick", 27}, {"ethan hunt", 32}}
	var jsonData4, err4 = json.Marshal(object)
	if err != nil {
		fmt.Println(err4.Error())
		return
	}

	var jsonString4 = string(jsonData4)
	fmt.Println(jsonString4)
}