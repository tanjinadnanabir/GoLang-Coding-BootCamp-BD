package main

import "fmt"

// // struct - a collection of fields and properties

// type Book struct{
// 	Title string
// 	Author string
// 	ISBN string
// 	Price float32
// 	Pages int
// }

func main(){

	// // Declaring this struct Book
	// var b1 Book

	// // assigning values
	// b1.Title = "An Introduction to Programming in Go"
	// b1.Author = "Coleb Dxsey"
	// b1.ISBN = "978-1478355823"
	// b1.Price = 10.5
	// b1.Pages = 140

	// // printing the values of struct Book
	// fmt.Println(b1)
	// fmt.Println(b1.Title)
	// fmt.Println(b1.Author)
	// fmt.Println(b1.ISBN)
	// fmt.Println(b1.Price)
	// fmt.Println(b1.Pages)

	// Alternative way to represent struct
	// struct literals
	// In computer science, a literal is a notation for representing a fixed value in source code
	// anonymous struct- An anonymous struct is just like a normal struct, but it’s defined without a name, and as such can’t be referenced elsewhere in code.

	// b1 := struct{
	// 	Title string
	// 	Author string
	// 	ISBN string
	// 	Price float32
	// 	Pages int
	// }{
	// 	Title: "An Introduction to Programming in Go",
	// 	Author: "Coleb Dxsey",
	// 	ISBN: "978-1478355823",
	// 	Price: 10.5,
	// 	Pages: 140,
	// }
	// fmt.Println(b1)

	var w1 Weight
	w1 = 30.5
	fmt.Println(w1)
}

