package main

import (
	"fmt"
	"net/url"
)

const myUrl = "http://localhost:3001/api/v1/patents/search?moudle=export-search"

func main() {
	fmt.Println("Handling URl in Go lang")

	fmt.Println("The Url is", myUrl)

	resultUrl, _ := url.Parse(myUrl)

	// fmt.Println(resultUrl.Scheme)
	// fmt.Println(resultUrl.Host)
	// fmt.Println(resultUrl.RawQuery)
	// fmt.Println(resultUrl.Port())

	qprams := resultUrl.Query()
	fmt.Printf("%T \n", qprams)
	fmt.Println(qprams)

	for key, val := range qprams {
		fmt.Println(key, " -> ", val)
	}

}
