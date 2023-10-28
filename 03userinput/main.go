package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user Input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the Rating for our Pizza:")

	// comma ok // err syntax

	input, _ := reader.ReadString('\n')

	fmt.Println("The Rating :", input)
	fmt.Printf("%T \n", input)
}
