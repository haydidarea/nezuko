package main

import "fmt"

func main() {
	var ujian = 80
	var absensi = 80

	var lulusujian = ujian >= 80
	var lulusabsensi = absensi >= 80

	var lulus = lulusujian && lulusabsensi

	fmt.Println(lulus)
}
