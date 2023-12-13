package main

import "fmt"

func main() {
	var names [3]string
	names[0] = "Apel"
	names[1] = "Blimbing"
	names[2] = "Ceri"

	fmt.Println(names[0])
	fmt.Println(names[1])
	fmt.Println(names[2])

	var values = [3]int{
		90,
		80,
		70}

	fmt.Println(values)
	
	fmt.Print(len(names))
	fmt.Println(len(values))

	var oke [12]string

	fmt.Println(len(oke))
}
