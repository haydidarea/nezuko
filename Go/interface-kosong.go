package main

import "fmt"

func oh(i int) interface{} {
	if i == 1 {
		return 1
	} else if i == 2 {
		return true
	} else {
		return "Salah "
	}
}

func main() {
	var data interface{} = oh(3) // nilai inputan
	fmt.Print(data)              //hasil
}
