package main

import "fmt"

func factorialloops(value int) int {
	hasil := 1 // =>kondisi berhenti
	for i := value; i > 0; i-- {
		hasil *= i // hasil =hasil x i
	}
	return hasil // loops fungsi
}
func recursive(value int) int {
	if value == 1 {
		return 1
	} else {
		return value * recursive(value-1) //recursi fungsi
	}
}

func main() {
	loop := factorialloops(4) //1x2x3x4  mengunakan fungsi loop atau pengulangan
	fmt.Println(loop)
	Recursive := recursive(5)
	fmt.Println(Recursive) // menggunakan fungsi recursive
}
