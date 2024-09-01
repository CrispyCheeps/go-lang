package main

import (
	"fmt"
	"os"
)

/*
	Arguments -> data argument opsionaal yg disisipkan ketika eksekusi program
	flag -> ekstensi dr argument
*/

func main() {
	var argsRaw = os.Args
	fmt.Printf("-> %#v\n", argsRaw)

	var args = argsRaw[1:]
	fmt.Printf("->%#v\n", args)
}