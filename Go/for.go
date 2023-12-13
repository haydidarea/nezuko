package main

import (
	"fmt"
)

func main() {
	counter := 1
	for counter <= 10 {
		fmt.Println("Perulangan ke -", counter)
		counter++
	}
	for counter := 11; counter <= 15; counter++ {
		fmt.Println("Perulangan ke -", counter) // for statment sederhana (simple nya)
	}
	slice := []string{"Admin", "User", "Eko"}

	for i := 0; i < len(slice); i++ {
		fmt.Println(slice[i]) // perulangan dengan range
	}
	for _, value := range slice {
		//fmt.Println("index", i, "=", value) // for range dengan index (no1) //jika no.1 pakai i
		fmt.Println(value) //pakai for _
	}

	person := make(map[string]string)
	person["name"] = "user"
	person["title"] = "Admin"

	for key, value := range person {
		fmt.Println(key, "=", value) //for range u/map

	}

}
