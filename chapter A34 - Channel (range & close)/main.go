package main

import (
	"fmt"
	"runtime"
)

/*
Channel direction

ch chan string  = channel bs "nerima" & "mengirim" data
ch chan <- string = hanya utk ngirim
ch <- chan string = hanya utk nerima
*/

//A34.1 Penerapan for - range - close
//berguna utk retreiving data dr banyak channel spy lebih mudah
func sendMessage(ch chan <- string) {
	for i := 0; i < 20; i++ {
		ch <- fmt.Sprintf("data %d ", i)
	}
	close(ch)
}
func printMessage(ch <- chan string) {
	for message := range ch {
		fmt.Println(message)
	}
}

func main() {
	runtime.GOMAXPROCS(2)

    // var messages = make(chan string)
    // go sendMessage(messages)
    // printMessage(messages)

	var messages = make(chan string)
	go sendMessage(messages)
	printMessage(messages)
}