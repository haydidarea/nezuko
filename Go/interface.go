package main

import "fmt"

type Animal struct {
	Name string
}

func (animal Animal) Getname() string {
	return animal.Name
}

/**type Hasname interface {
	GetName() string
}

func Sayhello(hasname Hasname) {
	fmt.Print("Hello ", hasname.GetName())

}

type Person struct {
	Name string
}

func (person Person) GetName() string {
	return person.Name
}*/

func main() {

	animal := Animal{Name: "kucing"}
	fmt.Println("sayhello", animal)
	/**var admin Person
	admin.Name = "Smith "
	Sayhello(admin)*/

}

//Struct adalah kumpulan definisi variabel
//(atau property) dan atau fungsi (atau method), yang dibungkus sebagai tipe data
//baru dengan nama tertentu.

/**interface adalah kumpulan definisi method yang tidak memiliki isi
(hanya definisi saja), yang dibungkus dengan nama tertentu.(bisa saling terhubung) interface biasa digunakan untuk struck*/
