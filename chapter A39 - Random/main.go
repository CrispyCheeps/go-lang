package main

import (
	"fmt"
	"math/rand"
	"time"
)

//A39.6 Random Tipe data String
var letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
var randomizer = rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

func randomString(length int) string {
	b := make([]rune, length)
	for i := range b {
		b[i] = letters[randomizer.Intn(len(letters))]
	}
	return string(b)
}

func main() {
	randomizer := rand.New(rand.NewSource(10))
	fmt.Println("random ke-1:", randomizer.Int())
	fmt.Println("random ke-2:", randomizer.Int())
	fmt.Println("random ke-3:", randomizer.Int())

	//A39.3 Unique seeds
	randomizer2 := rand.New(rand.NewSource(time.Now().UTC().UnixNano()))
	fmt.Println("random ke-1:", randomizer2.Int())
	fmt.Println("random ke-2:", randomizer2.Int())
	fmt.Println("random ke-3:", randomizer2.Int())
	fmt.Println("random ke-4:", randomizer2.Float32())
	fmt.Println("random ke-5:", randomizer2.Float64())
	fmt.Println("random ke-6:", randomizer2.Uint32())
	fmt.Println("random ke-7:", randomizer2.Uint64())

	//A39.5 Angka random index tertentu (dengan batas 0 hingga n-1)
	fmt.Println("random ke-8:", randomizer2.Intn(5))

	//A39.6 Random Tipe data String
	fmt.Println("random string 5 karakter: ", randomString(5))
}

	

