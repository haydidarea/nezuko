package main

import "fmt"

func GoodDay(name string) string {
	return "GoodDay " + name
}

func main() {
	Day := GoodDay
	result := (Day("Admin")) // ops function value
	fmt.Println(result)
}
