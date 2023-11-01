package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"coursename"`
	Price    int
	Platform string   `json:"website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("Json in Go")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	myCourses := []course{
		{"ReactJs", 399, "Google.com", "pas123", []string{"web", "react", "dev"}},
		{"Express", 499, "Google.com", "Test123", []string{"node", "backend", "dev"}},
		{"FastApi", 399, "Google.com", "zxy234", nil},
	}

	//package this data in Json data

	finalJson, err := json.MarshalIndent(myCourses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`
	{
		"coursename": "ReactJs",
		"Price": 399,
		"website": "Google.com",
		"tags": ["web","react","dev"]
    } 
	`)

	var myCourses course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json is Valid")
		json.Unmarshal(jsonDataFromWeb, &myCourses)
		fmt.Printf("My Courses %#v\n", myCourses)
	} else {
		fmt.Println("Not a valid json")
	}

	// some cases where you just want to add data to key value

	var myCustmdata map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myCustmdata)
	fmt.Printf("%#v\n", myCustmdata)

	for key, val := range myCustmdata {
		fmt.Println(key, val)
	}
}
