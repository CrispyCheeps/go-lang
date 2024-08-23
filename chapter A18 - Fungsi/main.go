package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

var randomizer = rand.New(rand.NewSource(time.Now().Unix()))

func main () {
	var names = []string{"john", "wick"}
	printMessage("Halo", names)

	var randomValue int
	randomValue = randomWithRange(2,10)
	fmt.Println("random number:", randomValue)

	randomValue = randomWithRange(2,10)
	fmt.Println("random number:", randomValue)

	randomValue = randomWithRange(2,10)
	fmt.Println("random number:", randomValue)
}

func printMessage(message string, arr []string) {
	var nameString = strings.Join(arr, " ")
	fmt.Println(message, nameString)
}

//A18.2 Fungsi dengan return value / nilai balik
func randomWithRange(min, max int) int {
	var value = randomizer.Int()%(max-min+1) + min
	return value
}

//A18.5 Deklarasi Parameter bertipe data sama
/*
	func nameOfFunc(param A type, paramB type, param C type) return Type
	func nameOfFunc(param A, param B, param C type) returnType
*/
func randommWithRange(min, max int) int

//Return sbg penghenti proses dlm fungsi
func divideNumber(m, n int) {
	if n==0 {
		fmt.Printf("invalid divider. %d cannot divided by %d\n", m, n)
		return
	}

	var res = m / n
	fmt.Printf("%d / %d = %d \n", m, n, res)
}