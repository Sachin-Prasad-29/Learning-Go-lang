package main

import "fmt"

func main() {
	fmt.Println("Maps in Go")

	lan := make(map[string]string)

	lan["JS"] = "JavaSccript"
	lan["RB"] = "RUby"
	lan["PY"] = "Python"
	lan["Ts"] = "Typescript"

	fmt.Println(lan)
	fmt.Println("Js", lan["JS"])

	// delete the value

	delete(lan, "RB")
	fmt.Println(lan)

	// loops in GO
	for key, value := range lan {
		fmt.Printf("For %v, Value is %v \n", key, value)
	}

}
