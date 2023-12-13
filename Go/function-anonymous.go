package main

import "fmt"

type blacklist func(string) bool

func pendaftaran(name string, blacklist blacklist) {
	if blacklist(name) == true {
		fmt.Println("Akses Di Block", name) // q1
	} else {
		fmt.Println("Welcome", name)
	}
}

//func blacklistuser(name string) bool {
//	return name == "user"
//}

//	func blacklistApple(name string) bool {
//		return name == "Appel" cara lama nya
//	}
func main() {

	blacklist := func(name string) bool {
		return name == "user"
	}
	pendaftaran("user", blacklist) //akses di block q1
	pendaftaran("Admin", blacklist)
}
