package main

import "fmt"

func main() {
	// single line comment- variable declaration
	var name string = "Coding Bootcamp"
	var my_name, email string = "Tanjin", "tanjin@mail.com"

	/*
	Multi-line comments
	*/

	fmt.Println(name, email, my_name)

	var c rune = 'A'
	var age int = 10
	var result float32 = 32.35

	fmt.Println(c, age, result)
}