package main

import (
	"fmt"
	"strings"
)

//A22.1 Penerapan fungsi sbg parameter
//A22.2 Alias skema closure
/*
	Nah kebetulan di skema closure filter di bawah ga terlalu pjg
	kalau skemanya cukup pjg, lebih baik pake alias; jadi dibuat type kyk gini:

	type FilterCallback func(string) bool

	func filter(data []string, callback FilterCallback) []string {
		....
	}
*/
func filter(data []string, callback func(string) bool ) []string {
	var result []string
	for _, each := range data {
		if filtered := callback(each); filtered {
			result = append(result, each)
		}
	}
	return result 
}

func main() {
	var data = []string{"wick", "jason", "ethan"}
	var dataContainsO = filter(data, func(each string) bool {
		/*
			string.Contains() berfungsi utk ngedeteksi apakah substring
			adalah bagian dr string; return value nya bool
		*/
		return strings.Contains(each, "o")
	})
	var dataLength5 = filter(data, func(each string) bool {
		return len(each) == 5
	})
	fmt.Println("data asli \t\t:", data)
	fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)
	fmt.Println("filter jumlah huruf \"5\"\t:", dataLength5)
	
}

//A22.2 Alias skema closure
