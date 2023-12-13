package main

import "fmt"

func logging() {
	fmt.Print("Selesai memanggil function ")
}

func runaps(value int) {
	defer logging() //defer function
	fmt.Print("Run aps ")
	hasil := 10 / value
	fmt.Print("hasil ", hasil)
	logging()
}

func main() {

	runaps(10) //jika value di buat 0 maka defer function tetap dijalankan
	// tetapi akan dihentikan oleh fungsi panic

}
