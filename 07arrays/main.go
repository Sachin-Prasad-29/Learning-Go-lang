package main

import "fmt"

func main() {
	fmt.Println("Array in Go ")
	var list [4]string
	list[0] = "sachin"
	// list[1] = "sameer"
	list[2] = "nitish"
	list[3] = "satya"

	fmt.Println("List is ", list)
	fmt.Println("List is ", len(list))

	var newlist = [8]string{"12", "15", "23"}

	fmt.Println("New list is ", newlist)
	fmt.Println("New list is ", len(newlist))

}
