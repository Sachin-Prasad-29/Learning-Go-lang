package main

import "fmt"

func main() {
	defer fmt.Println("First")
	defer fmt.Println("Second")
	fmt.Println("Hello")
	myDefer()
	defer fmt.Println("Third")
}

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
