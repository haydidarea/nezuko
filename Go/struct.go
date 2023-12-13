package main

import "fmt"

type Costumer struct {
	name, address string
	age           int // data awal struck
}

func main() {
	/**var Admin Costumer
	Admin.name = " Smith " // membuat isi data type struct ke-1
	Admin.address = "Europe "
	Admin.age = 25

	fmt.Print(Admin)
	fmt.Print(Admin.name)
	fmt.Print(Admin.address)
	fmt.Print(Admin.age)*/

	/**Admin := Costumer{
		name:    " Smith ",
		address: "Europe ", // membuat type data struct ke-2
		age:     25,
	}
	fmt.Print(Admin)*/

	Admin := Costumer{"Smith", "Europe", 25}
	fmt.Print(Admin) // membuat type data struck ke -3

}
