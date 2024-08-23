package main

import (
	"encoding/base64"
	"fmt"
)

//A46 Encode decode

/*
	Go menyediakan package encoding/base64, berisikan fungsi-fungsi untuk kebutuhan encode dan decode data
	 ke bentuk base64 dan sebaliknya. Data yang akan di-encode harus bertipe []byte, maka perlu dilakukan
	 casting untuk data-data yang tipenya belum []byte.

Proses encoding dan decoding bisa dilakukan via beberapa cara yang pada chapter ini kita akan pelajari.
*/

//A.46.1. Penerapan Fungsi EncodeToString() -> encode data dr string ke base64 & DecodeString()
func main() {
	var data = "john wick"

	var encodedString = base64.StdEncoding.EncodeToString([]byte(data))
	fmt.Println("encoded:", encodedString)

	var decodedByte, _ = base64.StdEncoding.DecodeString(encodedString)
	var decodedString = string(decodedByte)
	fmt.Println("decoded: ", decodedString)

	//A.46.2. Penerapan Fungsi Encode() & Decode()
	fmt.Println("a46.2")
	var encoded = make([]byte, base64.StdEncoding.EncodedLen(len(data)))
	base64.StdEncoding.Encode(encoded, []byte(data))
	var encodedString2 = string(encoded)
	fmt.Println(encodedString2)

	var decoded = make([]byte, base64.StdEncoding.DecodedLen(len(encoded)))
	var _, err = base64.StdEncoding.Decode(decoded, encoded)
	if err != nil {
		fmt.Println(err.Error())
	}
	var decodedString2 = string(decoded)
	fmt.Println(decodedString2)

	//A.46.3. Encode & Decode Data URL
	var data2 = "https://kalipare.com/"

	var encodedString3 = base64.URLEncoding.EncodeToString([]byte(data2))
	fmt.Println(encodedString3)

	var decodedByte2, _ = base64.URLEncoding.DecodeString(encodedString)
	var decodedString3 = string(decodedByte2)
	fmt.Println(decodedString3)
}
