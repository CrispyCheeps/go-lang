package main

import "fmt"

func main() {
	var point = 6

	switch {
	case point == 8:
		fmt.Println("Perfect")
	case (point < 8 ) && (point > 3) :
		fmt.Println("awesome")

		//efeknya case point < 5 dianggep always true
    fallthrough
case point < 5:
    fmt.Println("you need to learn more")
default:
    {
        fmt.Println("not bad")
        fmt.Println("you need to learn more")
    }
}
 	
}