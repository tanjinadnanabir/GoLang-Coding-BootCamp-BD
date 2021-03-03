package main

import "fmt"

// import "reflect"

func main(){

	// primitive data types
	// rune, byte
	// int, float32, string, bool

	// composite data types
	// // array

	// // array representation of golang: var variableName [size] dataType
	// var students [3] string
	// // fixed length
	// // same type
	// // index count start from 0

	// // assigning string values to the array students
	// students[0] = "Tanjin"
	// students[1] = "Yousuf"
	// students[2] = "Mohammad"

	// // printing specific value
	// fmt.Println(students[0])

	// // printing the array
	// // data pull, retreive or get
	// fmt.Println(students)	// it will show [   ]
	
	// // to show the length of array, use len() function
	// fmt.Println(len(students))

	// // array: short-hand way to assign
	// // string literals
	// products := [3] string{"iPhone", "Laptop", "Camera"}
	// fmt.Println(products)

	// // assign unlimited values
	// // implicit = ...
	// books := [...] string{"inception", "crisis", "perception", "friends"}
	// fmt.Println(books)

	// slice

	// assigning string values to the array students
	// students[0] = "Tanjin"
	// students[1] = "Yousuf"
	// students[2] = "Mohammad"

	// // slicing from 0 to before 2
	// x := students[0:2]	

	// alternative slice
	// x := make([] string, 0)

	// var fruits [] string
	// fruits = append(fruits, "Apple", "Lemon")
	// fmt.Println(fruits, len(fruits))

	// // to see the type of variable
	// fmt.Printf("%T\n", fruits)
	// fmt.Printf("%T", students)

	// a := reflect.TypeOf(students).Kind().String()
	// fmt.Println(a)

	// b := reflect.TypeOf(fruits).Kind().String()
	// fmt.Println(b)

	// map
	// var x map[string] string

	x := make(map[string]string)

	x["name"] = "Tanjin"
	x["age"] = "25"

	delete(x, "age")

	fmt.Println(x)
	fmt.Println(x["name"])
	// struct
}
