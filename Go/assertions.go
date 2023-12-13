package main

import "fmt"

func random() interface{} {
	return 2000 //data di sini harus mempunyai var penghubung yang sama denga di result string nya
}

func main() {

	var result interface{} = random()
	switch value := result.(type) {
	case string:
		fmt.Println("Value", value, "Is string") // cara aman menggunakan assertion adalah dengan menambakan fungsi switch
	case int:
		fmt.Println("Value", value, "Is int")
	default:
		fmt.Println("Nilai tidak di ketahui")
	}
	/**var resultString string = result.(string) //pastikan data sebelumnya adalah string atau sama jika tidak akan terjadi case panic
	fmt.Print(resultString)*/

}
