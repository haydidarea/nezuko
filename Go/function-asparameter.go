package main

import "fmt"

type Filter func(string) string

// menggunakan fungsi function alias
func Filters(name string, Filters Filter) {
	namefilter := Filters(name)
	fmt.Println("Hello", namefilter)
}
func spamFilter(name string) string {
	if name == "Apple" {
		return "***"
	} else {
		return name
	}
}

func main() {
	Filters("Admin", spamFilter)
	Filters("Apple", spamFilter)
}
