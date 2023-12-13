package main

import "fmt"

type Costumer struct {
	name, address string
	age           int // data awal struck
}

func (Costumer Costumer) Wlcm(name string) {
	fmt.Print("Ohayo ", name, "Kak ", Costumer.name)

}

func main() {
	var Admin Costumer
	Admin.name = " Smith " // membuat isi data type struct ke-1
	Admin.address = "Europe "
	Admin.age = 25

	Admin.Wlcm("user ")
}
