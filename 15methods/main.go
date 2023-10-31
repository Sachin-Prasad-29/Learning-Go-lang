package main

import "fmt"

func main() {
	fmt.Println("Struct in Go")

	// no inferitance in golang; no super or parent

	user := User{"Sachin", "sachin@gmail.com", true, 25}

	// fmt.Println(user.Name)
	// fmt.Println(user.Email)
	// fmt.Printf("%+v\n", user)
	// fmt.Printf("%v\n", user.Name)
	// fmt.Println(user.Age)

	user.GetStatus()
	user.GetDetails()
	user.NewMail()
	user.GetDetails()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User active :", u.Status)
}

func (u User) GetDetails() {
	fmt.Println("Name : ", u.Name)
	fmt.Println("Email : ", u.Email)
	fmt.Println("Age : ", u.Age)
}

func (u User) NewMail() {
	u.Email = "test@go.dev"
	fmt.Println("Email of this user is:", u.Email)
}
