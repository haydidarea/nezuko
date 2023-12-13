package main

import "fmt"

func nameval() (A1 string, A2 string, A3 string) {
	A1 = "Bima sakti"
	A2 = "klas"
	A3 = "12"

	return
}
func main() {
	A1, A2, A3 := nameval()
	fmt.Println(A1, A2, A3)
}
