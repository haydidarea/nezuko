package main

import "fmt"

func main() {
	type NoKTP string
	type Single bool

	var noKTPuser NoKTP = "12345678910"
	var Singlestatus Single = true

	fmt.Println(noKTPuser)
	fmt.Println(Singlestatus)
}
