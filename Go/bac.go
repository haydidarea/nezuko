package main

import "fmt"

func main() {
	for i := 0; i <= 10; i++ {
		if i == 7 {
			break // if break count akan stop di value break
		}
		fmt.Println("For -", i)
	}
	for i := 90; i <= 100; i++ {
		if i%2 == 1 {
			continue // if continue value break akan di skip atau bisa di modulo %2 dibagi
			// maka bil genap akan di hide jika value 0 tetapi jika value 1 akan sebaliknya
		}
		fmt.Println("For -", i)
	}
}
