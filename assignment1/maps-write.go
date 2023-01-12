package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	map1 := map[string]string{
		"a": "apple",
		"b": "banana",
	}

	jsonStr, err := json.Marshal(map1)
	if err != nil {
		fmt.Printf("Error: %s", err.Error())
	}
	err = ioutil.WriteFile("userfile.json", jsonStr, 0644)
	if err != nil {
		log.Fatal(err)
	}
	jsonStr, err = ioutil.ReadFile("userfile.json")
	if err != nil {
		log.Fatal(err)
	}

}
