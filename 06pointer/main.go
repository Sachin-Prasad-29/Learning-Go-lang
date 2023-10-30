package main

import "fmt"

func main() {
	fmt.Println("Pointer in Go")

	// var ptr *int

	// fmt.Println("Value of Ptr ", ptr)

	myNumber := 23

	var ptr = &myNumber

	fmt.Println("Value of Ptr", ptr)  // memory location is printed
	fmt.Println("Value of Ptr", *ptr) // actual value is printed

	*ptr = *ptr + 2
	fmt.Println("New Value is: ", myNumber)
}
