package main

import (
	"fmt"
	"time"
)

//A42 Time Duration
/*
	time.Duration() sgt berguna utk benchmarking atau operasi2 yg membutuhkan informasi durasi waktu.
*/

func main() {
	start := time.Now()

	time.Sleep(5 * time.Second)

	duration := time.Since(start)

	fmt.Println("time elapsed in seconds:", duration.Seconds())
	fmt.Println("time elapsed in minutes:", duration.Minutes())
	fmt.Println("time elapsed in hours:", duration.Hours())

	//A42.4 Kalkulasi durasi antara 2 objek waktu
	t1 := time.Now()
	time.Sleep(5 * time.Second)
	t2 := time.Now()

	durationn := t2.Sub(t1)

	fmt.Println("time elapsed in seconds:", durationn.Seconds())
	fmt.Println("time elapsed in minutes:", durationn.Minutes())
	fmt.Println("time elapsed in hours:", durationn.Hours())
}