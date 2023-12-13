package main

import "fmt"

func main() {
	Person := map[string]string{
		"name":   "User",
		"addres": "none",
	}

	Person["title"] = "Progamer" // menambahkan atau edit

	fmt.Println(Person)
	fmt.Println(Person["name"])
	fmt.Println(Person["addres"])
	fmt.Println(Person["title"])

	var book map[string]string = make(map[string]string)
	book["title"] = "Belajar"
	book["author"] = "user"
	book["ups"] = "salah"
	fmt.Println(book)
	fmt.Println(len(book))

	delete(book, "ups")
	fmt.Println(book)
	fmt.Println(len(book))

}
