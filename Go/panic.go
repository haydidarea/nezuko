package main

import "fmt"

func endapp() {
	message := recover() // Menempatkan fungsi recover
	if message != nil {
		fmt.Print("Error dengan message") //gunakan fungsi if nill jika tidak sama dengan error
	}
	fmt.Print("Done... ")
}

func runapp(error bool) {
	defer endapp() // awali dengan function defer
	if error {
		panic("Error ") // baru panic function di panggil
	}

	fmt.Print("Aplikasi berjalan ")
}

func main() {
	runapp(true)
	fmt.Print("Admin ")

}
