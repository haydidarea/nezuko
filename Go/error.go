package main

import (
	"errors"
	"fmt"
)

func Pembagian(nilai int, pembagi int) (int, error) {
	if pembagi == 0 {
		return 0, errors.New("Pembagi tidak boleh bilangan kosong")
	} else {
		result := nilai / pembagi
		return result, nil
	}

}

func main() {
	hasil, err := Pembagian(100, 10)

	if err == nil {
		fmt.Print("Hasil ", hasil) // jika tidak kosong maka cetak hasil

	} else {
		fmt.Print("Error", err.Error()) //jika kosong tampilkan error
	}

}
