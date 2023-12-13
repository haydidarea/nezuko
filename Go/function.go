package main

import "fmt"

func sayhello() {
	fmt.Println("Hello World")
}

func main() {
	for i := 0; i < 10; i++ {
		sayhello()
	}
}
