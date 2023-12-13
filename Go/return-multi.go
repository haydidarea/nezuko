package main

import "fmt"

func namalengkap() (string, string, string) {
	return "user ", "Admin ", "New"
}

func main() {
	A1, _, A3 := namalengkap() // di deklarasikan A1 dan A2 di mengantikan nama awak dan akhir
	fmt.Println(A1)
	//fmt.Println(A2) A2 di ignore(di skip)
	fmt.Println(A3)
}
