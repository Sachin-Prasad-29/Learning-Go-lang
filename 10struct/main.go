package main

import "fmt"

func main() {
	fmt.Println("Struct in Go")

	// no inferitance in golang; no super or parent

	user := User{"Sachin", "sachin@gmail.com", true, 25}

	fmt.Println(user.Name)
	fmt.Println(user.Email)
	fmt.Printf("%+v\n", user)
	fmt.Printf("%v\n", user.Name)
	fmt.Println(user.Age)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
