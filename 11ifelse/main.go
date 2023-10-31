package main

import "fmt"

func main() {
	fmt.Println("If else In Go")

	loginCount := 1

	var result string
	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch Out"
	} else {
		result = "Equal to Ten"
	}

	fmt.Println(result)

	if num := 3; num < 10 {
		fmt.Println("Num is less than 10")
	} else {
		fmt.Println("Num is Not Less than 10")
	}

	// if err != nill{

	// }
}
