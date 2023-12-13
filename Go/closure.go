package main

import "fmt"

func main() {
	name := "user"
	counter := 0 //data awal

	increment := func() {
		name := "Admin" //=> gunakan := supaya terdeklarasi
		fmt.Println("increment")
		fmt.Print(name) //mencetak hasil dari name :=
		counter++ // tanpa sengaja menambah counter++ maka data counter di atas nilainya berubah
		//=> jadi kekita kita megubah data di dalam kurung kurawal ini maka
		//secara otomatis data di kurung kurawal sebelumnya akan akan berubah

	}
	increment()
	increment()
	fmt.Println(counter)
	fmt.Println(name)
}
