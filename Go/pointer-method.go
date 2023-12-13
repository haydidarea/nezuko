package main

import "fmt"

type Man struct {
	Name string
}

func (man *Man) married() {
	man.Name = "Mr." + man.Name     // + kan fungsi * di parameter supaya menjadi pointer
	fmt.Println("Method", man.Name) // untuk mencek pass by value or duplikat nya apakah sudah di awali dengan kata penguhung nya?
}

func main() {
	Smith := Man{"Smith"}
	Smith.married() //harapan nya ketika kita call marrie akan mengubah nama awalan dari subject
	// dan saat memakai pointer method data di main func tidak perlu di rubah sama sekali atau menanbahkan fungsi &
	fmt.Println(Smith) // akan menghasilkan poiter tanpa method
}
