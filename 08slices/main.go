package main

import "fmt"

func main() {
	fmt.Println("Slices in Go")

	var list = []string{"sachin", "sameer", "satya"}
	fmt.Printf("Type of list %T\n", list)

	list = append(list, "nitish", "karan")

	fmt.Println("List is", list)

	list = append(list[1:3])
	fmt.Println("List is", list)

}
