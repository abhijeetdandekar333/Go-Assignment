package main

import (
	"encoding/json"
	"fmt"
)

// sample json object to be parsed
var companyJSON = `{
	"name" : "GOLinuxCloud",
        "years_of_service" : "5",
        "nature_of_company" : "Online Academy",
        "no_of_staff" : "10"
}`

func main() {

	// Declared an empty map interface
	var dataResult map[string]interface{}

	// Unmarshal the JSON to the interface.its same as decode
	err := json.Unmarshal([]byte(companyJSON), &dataResult)
	if err != nil {
		fmt.Println(err)
	}

	// Print elements in map on the terminal the key and its value
	for key, value := range dataResult {
		fmt.Printf("%s : %v \n", key, value)
	}
}
