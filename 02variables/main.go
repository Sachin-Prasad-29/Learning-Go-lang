package main

import (
	"fmt"
)

// declaring constant

const LoginToken string = "Contanst sachin" //Public contant variables because of Captial first letter

func main() {
	var username string = "Sachin"
	fmt.Println(username)
	fmt.Printf("%T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("%T \n", isLoggedIn)

	var smallVal uint8 = 255
	fmt.Println(smallVal)
	fmt.Printf("%T \n", smallVal)

	var smallFloat float64 = 255.234234523523523
	fmt.Println(smallFloat)
	fmt.Printf("%T \n", smallFloat)

	// default values and some aliases

	var anotherVariable int
	fmt.Println(anotherVariable)
	var anotherString string
	fmt.Println(anotherString)

	//implicit type no need to provide the variable type

	var website = "sachinis here"
	fmt.Println(website)

	// no Var style
	numberOfUsers := 23838.3
	fmt.Println(numberOfUsers)

	fmt.Println(LoginToken)

}
