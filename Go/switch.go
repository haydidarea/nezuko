package main

import "fmt"

func main() {

	name := "Admin"
	switch name {

	case "user":
		fmt.Println("hai,user")

	case "Admin":
		fmt.Println("hai Admin")

	default:
		fmt.Println("Tidak dikenali")
	}

	switch length := len(name); length > 5 {
	case true:
		fmt.Println("Too long")
	case false:
		fmt.Println("Okay")

	}
	length := len(name) // if versi sederhana
	switch {
	case length > 10:
		fmt.Println("Panjang sekali")
	case length > 4:
		fmt.Println("Cukup panjang")
	default:
		fmt.Println("kosong")

	}
}
