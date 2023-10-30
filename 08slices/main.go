package main

import (
	"fmt"
)

func main() {
	fmt.Println("Slices in Go")

	var list = []string{"sachin", "sameer", "satya"}
	// fmt.Printf("Type of list %T\n", list)

	list = append(list, "nitish", "karan")

	// fmt.Println("List is", list)

	list = append(list[1:3])
	fmt.Println("List is", list)

	// make keyword to make slices

	highScores := make([]int, 4)

	highScores[0] = 123
	highScores[1] = 823
	highScores[2] = 423
	highScores[3] = 323

	// highScores = append(highScores, 555, 444, 33313)

	// fmt.Println(highScores)

	// sort.Ints(highScores)

	// fmt.Println(highScores)
	// fmt.Println(sort.IntsAreSorted(highScores))

	// how to remove a value from slices based on index

	var courses = []string{"reactjs", "java", "python"}

	fmt.Println(courses)
	var index int = 1
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)

}
