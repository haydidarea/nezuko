package main

import "fmt"

func sayhelloto(fristname string, lastname string) {
	fmt.Println("ohayo", fristname, lastname) //2 parameter
}

func main() {
	fristname := "dodo" // bisa di call dengan nama parameter nya
	sayhelloto("fristname", "lastname")
}
