package main

import "fmt"

func Newmap(name string) map[string]string {
	if name == "" {
		return nil
	} else {
		return map[string]string{
			"name": name,
		}

	}
}
func main() {
	//person := Newmap("Smith") dengan nill
	var persom map[string]string = Newmap("Smith")
	if persom == nil {
		fmt.Print("Data Kosong ")
	} else {
		fmt.Println(persom)
	}
}
