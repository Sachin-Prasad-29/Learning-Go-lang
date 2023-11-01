package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	fmt.Println("Web server request in Go")
	// PerformGetRequest()
	// PerformPostRequest()
	PerformFormPostRequest()
}

func PerformGetRequest() {
	const myurl = "http://localhost:5000/get"

	response, err := http.Get(myurl)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("Status code: ", response.StatusCode)
	fmt.Println("Content length is: ", response.ContentLength)
	var responseString strings.Builder
	content, _ := ioutil.ReadAll(response.Body)
	responseString.Write(content)

	fmt.Println(responseString.String())
	// fmt.Println(string(content))
}

func PerformPostRequest() {
	const myurl = "http://localhost:5000/post"

	//fake json payload

	requestBody := strings.NewReader(`
		{
			"courseName":"Let's go with Go Lang",
			"price":13,
			"platform":"Google"
		}
	`)

	response, err := http.Post(myurl, "application/json", requestBody)

	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	constent, _ := ioutil.ReadAll(response.Body)
	fmt.Println(constent)
	fmt.Println(string(constent))
}

func PerformFormPostRequest() {
	const myurl = "http://localhost:5000/postform"

	//formdata
	data := url.Values{}

	data.Add("name", "sachin")
	data.Add("age", "23")
	data.Add("email", "sakok@gmai.com")

	response, err := http.PostForm(myurl, data)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	content, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(content))

}
