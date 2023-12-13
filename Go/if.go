package main

import "fmt"

func main() {

	var name = "Admin" // variable perbandingan nilai if

	if name == "Admin" {
		fmt.Println("Benar") //cek kondisi pertama

	} else if name == "User" {
		fmt.Println("Okay") //cek kondisi ke-2,else if bisa di +kan sesuai kebutuhan
	} else {
		fmt.Println("Salah") // jika tidak ditemukan nilai true di kondisi 1/2 maka salah hasilnya
	}

	if length := len(name); length > 5 {
		fmt.Println("Too long") //deklarasi simple
	} else {
		fmt.Println("Okay")
	}
}
