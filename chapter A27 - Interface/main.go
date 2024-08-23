package main

import (
	"fmt"
	"math"
)

/*
	interface adalah definisi suatu kumpulan method yg tdk memiliki isi.
	Jd cuma definisi header/schema-nya aja.
*/

type hitung interface {
	luas() float64
	keliling() float64
}

//Objek / struct lingkaran
type lingkaran struct {
	diameter float64
}

func(l lingkaran) jariJari() float64 {
	return l.diameter / 2
}

func(l lingkaran) luas() float64 {
	return math.Pi * math.Pow(l.jariJari(), 2)
}

func(l lingkaran) keliling() float64 {
	return math.Pi * l.diameter
}


//objek / struct persegi
type persegi struct {
	sisi float64
}

func (p persegi) luas() float64 {
	return math.Pow(p.sisi, 2)
}

func (p persegi) keliling() float64 {
	return p.sisi * 4
}

//A27.2 Embedded interface 

type hitung2d interface {
	luas() float64
	keliling() float64
}

type hitung3d interface {
	volume() float64
}

type count interface{
	hitung2d
	hitung3d
}

type kubus struct {
	sisi float64
}

func (k *kubus) volume() float64 {
	return math.Pow(k.sisi, 3)
}

func (k *kubus) luas() float64 {
	return math.Pow(k.sisi, 2) * 6
}

func (k *kubus) keliling() float64 {
	return k.sisi * 12
}

func main() {
	var bangunDatar hitung
	
	bangunDatar = persegi{10.0}
	fmt.Println("======= persegi")
	fmt.Println("luas : ", bangunDatar.luas())
	fmt.Println("keliling: ", bangunDatar.keliling())

	bangunDatar = lingkaran{14.0}
	fmt.Println("======= lingkaran")
	fmt.Println("luas : ", bangunDatar.luas())
	fmt.Println("keliling: ", bangunDatar.keliling())
	//(lingkaran) ini adalah cara casting di go pd objek interface spy bs diakses
	//nama metode casting kek gini namanya type assertion
	fmt.Println("keliling: ", bangunDatar.(lingkaran).jariJari())

	//contoh casting
	var bangunDatar2 hitung = lingkaran{14.0}
	var bangunLingkaran lingkaran = bangunDatar2.(lingkaran)

	bangunLingkaran.jariJari()

	//implementasi A27.2 Embedded interface 
	fmt.Print("\n\n")
	fmt.Println("//implementasi A27.2 Embedded interface ")
	var bangunRuang count = &kubus{4}

	fmt.Println("====kubus")
	fmt.Println("luas :", bangunRuang.luas())
	fmt.Println("keliling :", bangunRuang.keliling())
	fmt.Println("volume :", bangunRuang.volume())
}