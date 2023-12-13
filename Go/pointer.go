package main

import "fmt"

type Addres struct {
	city, province, country string
}

func Changecountrytoindonesia(addres *Addres) {
	addres.country = "Indonesia" //pointer di function jadikan Addres seabgai pointer dengan menambahkan*
}

func main() {
	/**addres1 := Addres{"Kendal", "Jawa tengah", "Indonesia"}
	addres2 := &addres1 // tambahkan ops & supaya data addres 2 sama dengan deklarasi addres 1*/

	var addres1 Addres = Addres{"Kendal", "Jawa tengah", "Indonesia"}
	var addres2 *Addres = &addres1    //jika dijabarkan lebih efektif pakai yang atas
	var addres3 *Addres = &addres1    //untuk mencek hasil
	var addres4 *Addres = new(Addres) //mengunakan function new di dalam pointer akan menghasilkan data kosong
	addres4.city = "kebumen"          //maka data addres4 akan berubah

	var alamat = Addres{
		city:     "Demak",
		province: "Jawa tengah",
		country:  "",
	}

	addres2.city = "Boja"
	*addres2 = Addres{"malang", "Jawa timur", "Indonesia"}
	//jika di tambhkan ini maka data addres2 akan berubah
	//dan + kan * di awal pointer yang di tuju supaya data yang di refrensikan sama dengan data di pointernya

	fmt.Println(addres1) //gunakan ctrl + spasi untuk mengetahui bila suatu addres ada * nya maka di sebut sebagai pointer
	fmt.Println(addres2)
	fmt.Print(addres3)
	fmt.Println(addres4)

	var alamatpointer *Addres = &alamat     // atau bisa jugda menggunakan fungsi ini jika sudah mempunyai pointer
	Changecountrytoindonesia(alamatpointer) //pointer di function + & supaya jadi pointer
	fmt.Println(alamat)
}
