package main

import "fmt"

func main() {
	//A17.1 Penggunaan Map

	var chicken map[string]int
	chicken = map[string]int{}

	chicken["januari"] = 50
	chicken["februari"] = 40

	fmt.Println("januari", chicken["januari"]) // januari 50
	fmt.Println("Mei", chicken["Mei"]) // mei 0

	//A17.2 Inisialisasi Nilai Map
	var data map[string]int
	data = map[string]int{}
	data["one"] = 1
	fmt.Println(data["one"])

	//cara horizontal
	var chicken1 = map[string]int{"januari" : 50, "feburari":40}

	//cara vertikal
	var chicken2 = map[string]int{
		"januari" : 50,
		"februari" : 40,
	}

	fmt.Println(chicken1)
	fmt.Println(chicken2)

	//map bs di-inisialisasi tanpa nilai awal. Syaratnya pake tanda kurung kerawal atau...
	var _ = map[string]int{}
	var _ = make(map[string]int)
	var _ = *new(map[string]int)

	//A17.3 Iterasi item Map pake for-range
	fmt.Println("A17.3")

	var chickenn = map[string]int{
		"januari" : 50,
		"februari" : 40,
		"maret" : 34,
		"april" : 67,
	}

	for key, val := range chickenn {
		fmt.Println(key, " \t:", val)
	}

	//A17.4 Menghapus item map
	fmt.Printf("\n\n")
	fmt.Printf("A17.4\n")
	fmt.Println(len(chickenn))
	fmt.Println(chickenn)
	delete(chickenn, "januari")
	fmt.Println(len(chickenn))
	fmt.Println(chickenn)

	//A17.5 Deteksi keberadaan item dg key tertentu
	var value, isExist = chicken["mei"]

	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("Item is not exists")
	}

	//A17.6 Kombinasi Slice & Map
	/*
		penerapannya contohnya pake []map[string]int

		berarti kyk gitu itu adalah sebuah slice yg stiap elemennya adalah map[string]int
		map yg tipe data nya int yg keynya adalah string 
	*/
	fmt.Printf("\n\n")
	fmt.Println("A17.6")
	var ayams = []map[string]string{
		{"name" : "chicken blue", "gender" : "male"},
		{"name" : "chicken red", "gender" : "male"},
		{"name" : "chicken yellow", "gender" : "female"},
	}

	for _, ayam := range ayams {
		fmt.Println(ayam["gender"], ayam["name"])
	}

	//Dalam []map[string]string bs aj tiap elemen punya key yg beda2

	var datas = []map[string]string{
		{"name" : "chicken kampus", "gender" : "male", "color" : "brown"},
		{"address" : "mangga street", "id" : "k001"},
		{"community" : "chicken lovers"},
	}

	for _, data := range datas{
		fmt.Println(data["id"], data["name"])
	}

	

}