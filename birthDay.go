package main
import "fmt" 

func main(){
 
	var size int = 15  

	fmt.Println(" ")

	for a := size/2; a <= size; a = a+2{ 

		// FOR SPACE BEFORE PEAK-1 : PART 1 
		for b := 1; b < size-a; b = b+2{
			fmt.Print(" ")
		}

		// FOR PRINTING PEAK-1 : PART 2 
		for b := 1; b <= a; b++{
			fmt.Print("*") 
		}
	
		// FOR SPACE B/W PEAK-1 AND PEAK-2 : PART 3 
		for b := 1; b <= size-a; b++ {
			fmt.Print(" ")
		}
		
		// FOR PRINTING PEAK-2 : PART 4 
		for b := 1; b <= a-1; b++{
			fmt.Print("*")
		}

		fmt.Println(" ") 
	} 

	/*FOR THE BASE OF HEART ie. THE INVERTED TRIANGLE */

	for a := size; a >= 0; a-- { 
		// FOR SPACE BEFORE THE INVERTED TRIANGLE : PART 5 
		for b := a; b < size; b++{
			fmt.Print(" ") 
		}

		// FOR PRINTING THE BASE OF TRIANGLE : PART 6 
		for b := 1; b <= ((a * 2) - 1); b++ {
			fmt.Print("*")
		}

		fmt.Println(" ")
	}

	fmt.Println(" Happy Birthday Mostain Sir! ")
	fmt.Println(" \n Wishing you a calm and relaxed birthday this year! ")
}
