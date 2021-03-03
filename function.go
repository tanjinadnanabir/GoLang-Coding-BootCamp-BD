package main
import "fmt"

/*// function named add to add two number
// 1st way
func add(x int, y int) int{
	//body
	r := x+y
	return r
}

// 2nd way: for same type of parameters
func add(x, y int) int{
	//body
	r := x+y
	return r
}

// 3rd way define (r int)
func add(x int, y int) (r int){
	//body
	r = x+y
	return r
}

// 4th way
func add(x int, y int) (r int){
	//body
	r = x+y
	return
}*/

/*func rectangle(l int, b int) (area int, parameter int) {
	parameter = 2 * (l + b)
	area = l * b
	return // Return statement without specify variable name
}*/

/*func update(a *int, t *string) {
	*a = *a + 5      // defrencing pointer address
	*t = *t + " Doe" // defrencing pointer address
	return
}*/

func main(){
	// calling the function
	// x:= add(10, 13)
	// a, p := rectangle(10, 10)
	// fmt.Println(a, p)
	/*a := 10
	t := "root"
	update(&a, &t)
	fmt.Println(a, t)*/

	// anonymous function
	/*a := func(x, y int)(r int){
		r = x*y
		return
	}
	fmt.Println(a(10,10))*/

	a := func(x, y int)(r int){
		r = x*y
		return
	}(10,10)

	fmt.Println(a)
}

