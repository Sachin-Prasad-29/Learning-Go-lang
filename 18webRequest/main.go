package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "https://jsonplaceholder.typicode.com/posts"

func main() {
	fmt.Println("Web Request")

	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Type of response is %T \n", response)
	// fmt.Print(response)
	defer response.Body.Close()

	dataBytes, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(dataBytes))
}
