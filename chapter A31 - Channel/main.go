package main

import (
	"fmt"
	"runtime"
)


func main() {
	// Channel berbentuk variabel
	// dibuat menggunakan kombinasi keyword "make" dan "chan"
	/*Utk menghubungkan goroutine satu dg GR lain dg mekanisme serah terima
	Jadi, channel ini dipakai sbg media perantara pengirim data dan juga penerima data*/

	runtime.GOMAXPROCS(2)

	var messages = make(chan string)

	var sayHelloTo = func(who string) {
		var data = fmt.Sprintf("hello %s", who)
		messages <- data
	}

	go sayHelloTo("john wick")
	go sayHelloTo("ethan hunt")
	go sayHelloTo("jason bourne")

	var message1 = <- messages
	fmt.Println(message1)

	var message2 = <- messages
	fmt.Println(message2)

	var message3 = <- messages
	fmt.Println(message3)

	//A31.2 Channel sbg tipe data parameter
	for _, each := range []string{"wick", "hunt", "bourne"} {
		go func(who string) {
			var data = fmt.Sprintf("Hello %s", who)
			messages <- data
		} (each)
	}
	for i := 0;i < 3; i++ {
		printMessage(messages)
	}
}

//A31.2 Channel sbg tipe data parameter
func printMessage(what chan string) {
	fmt.Println(<-what)
}