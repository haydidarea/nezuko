package main

import "fmt"

func wlcm(name string) string {
	if name == "" {
		return "Nama kosong"
	} else {
		return "Konichiwa " + name
	}
}

func main() {
	D1 := wlcm("")
	fmt.Println(D1) // secara deklarasi

	fmt.Println(wlcm("Admin")) //cara ringkas nya
}
