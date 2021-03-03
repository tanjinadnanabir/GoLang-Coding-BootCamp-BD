package main

import "fmt"

// package level scope
var w1 Weight

func main(){
	// function level scope
	// var w1 Weight
	w1 = 30.5

	fmt.Println(w1, name)
}