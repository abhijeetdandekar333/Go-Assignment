package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string `json:"CourseName"`
	Price    int
	Platform string   `json:"Website"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty"`
}

func main() {
	fmt.Println("hello")
	// EncodeJson()
	DecodeJson()
}

func EncodeJson() {
	lcoCourses := []course{
		{"ReactJS Bootcamp", 299, "Learncode", "abc123", []string{"web-dev", "js"}},
		{"Python Bootcamp", 199, "Learncode", "abc123", []string{"web-dev", "js"}},
		{"Web Bootcamp", 499, "Learncode", "abc123", nil},
	}

	finalJson, err := json.MarshalIndent(lcoCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalJson))

}

func DecodeJson() {
	jsonDataFromWeb := []byte(`[
        {
                "CourseName": "ReactJS Bootcamp",
                "Price": 299,
                "Website": "Learncode",
                "tags": [
                        "web-dev",
                        "js"
                ]
    
}]
`)

	var lcocourse course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("Json is valid")
		json.Unmarshal(jsonDataFromWeb, &lcocourse)
		fmt.Printf("%#v\n", lcocourse)
	} else {
		fmt.Println("Json is not valid")
	}

}
