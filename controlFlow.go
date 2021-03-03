package main
import "fmt"

func main(){
	/*fmt.Print("Enter your age: ")
	var age int
	fmt.Scanf("%d",&age)*/

	//if boolean_expression{
		//logic statement
	//}

	/*if age < 3 {	// 0 to 2
		fmt.Println("Infant")
	} else if age >= 3 && age < 13 {	// 3 to 12
		fmt.Println("Child")
	} else if age >= 13 && age <= 18 {	// 13 to 18
		fmt.Println("Teenage")
	} else {	// > 18
		fmt.Println("Adult")
	}*/

	/*switch age {
	case 1,2:
		fmt.Println("Infant")
		fallthrough
	case 3,4,5,6,7,8,9,10,11,12:
		fmt.Println("Child")
	case 13,14,15,16,17,18:
		fmt.Println("Teenage")
	default:
		fmt.Println("Adult")
	}*/

	// for loop
	// for initialization;condition;increment
	// i++ means i+1
	/*for i:=0; i<=9; i++ {
		fmt.Println(i)
		// will print 0 to 9
	}*/

	// for range
	/*entry := []string{"Jack","John","Jones"}
	for i, val := range entry {
	  fmt.Printf("At position %d, the character %s is present\n", i, val)
	}*/

	/*students := []string{"Jack","John","Jones"}
	for i, std := range students {
		fmt.Println(i, std)
	}*/

	// for boolean_expression
	/*i := 0
	for i<10 {
		fmt.Println(i, "Hello!")
		i++
	}*/
}