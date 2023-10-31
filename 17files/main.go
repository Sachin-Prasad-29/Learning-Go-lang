package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Files in Go")

	content := "This is some Content"

	file, err := os.Create("./newfile.txt")

	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("Length is : ", length)

	defer file.Close()
	readFile("./newfile.txt")

}

func readFile(filename string) {
	byteData, err := ioutil.ReadFile(filename)

	if err != nil {
		panic(err)
	}
	fmt.Println("The data of file \n", string(byteData))

}
