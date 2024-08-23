package main

import (
	"fmt"
	"regexp"
)

func main() {
	var text = "banana burger soup"
	var regex, err = regexp.Compile(`[a-z]+`)

	if err != nil{
		fmt.Println(err.Error())
	}

	//Method FindAllString
	var res1 = regex.FindAllString(text,2)
	fmt.Printf("%#v\n", res1)

	var res2 = regex.FindAllString(text, -1)
	fmt.Printf("%#v\n", res2)

	//Method MatchString()
	var isMatch = regex.MatchString(text)
	fmt.Println(isMatch)

	//Method FindString()
	var str = regex.FindString(text)
	fmt.Println(str)

	//Method FindStringIndex
	//Digunakan utk mencari index string kembalian hasil dr operasi regexp

	var idx = regex.FindStringIndex(text)
	fmt.Println(idx)	
	var str2 = text[0:6]
	fmt.Println(str2)
	// "banana"

	//A.45.5. Method FindAllString()
	var str3 = regex.FindAllString(text, -1)
	fmt.Println(str3)

	var str4 = regex.FindAllString(text,1)
	fmt.Println(str4)

	//A.45.6. Method ReplaceAllString()
	var str5 = regex.ReplaceAllString(text, "potato")
	fmt.Println(str5)

	//A.45.7. Method ReplaceAllStringFunc()
	var str6 = regex.ReplaceAllStringFunc(text, func(each string) string {
		if each == "burger" {
			return "potato"
		}
		return each
	}) 
	fmt.Println(str6)

	//A.45.8. Method Split()
	var str7 = regex.Split(text, -1)
	fmt.Printf("%#v\n", str7)
}