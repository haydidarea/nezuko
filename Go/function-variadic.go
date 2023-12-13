package main

import "fmt"

func Sum(number ...int) int {
	total := 0
	for _, value := range number {
		total += value
	}
	return total
}

func main() {
	total := Sum(2, 7) // dengan ada nya titik 3 didalam kurung
	// mendefinisikan argumen variadic dengan type data nya
	fmt.Println(total)
	slice := []int{1, 2, 3, 4, 5}
	total = Sum(slice...) //gunakan titik 3 supaya menjadi variable argumen
	fmt.Println(total)
}

//...int dinamic parameter --> akan di anggap sebagai array
//[]int array
