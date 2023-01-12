// package main

// import (
// 	"encoding/json"
// 	"fmt"
// 	"io/ioutil"
// )

// type User struct {
// 	Name        string
// 	Age         int
// 	Designation string
// 	Grade       float32
// }

// func main() {

// 	user1 := User{}
// 	user1.Name = "abhijeet"
// 	user1.Age = 23
// 	user1.Designation = "Software Engineer"
// 	user1.Grade = 3.0

// 	content, err := json.Marshal(user1)
// 	if err != nil {
// 		panic(err)
// 	}
// 	err = ioutil.WriteFile("practise.json", content, 0644)
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(err)
// 	//reading file
// 	databyte, err := ioutil.ReadFile("practise.json")
// 	if err != nil {
// 		panic(err)
// 	}
// 	fmt.Println(string(databyte))
// 	user2 := User{}
// 	err = json.Unmarshal(content, &user2)
// 	fmt.Println(user2.Name)

// }
package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type User struct {
	Name        string
	Age         int
	Designation string
	Grade       float32
}

func main() {
	var user1 User
	user1.Name = "abhijeet"
	user1.Age = 23
	user1.Designation = "Software Engineer"
	user1.Grade = 3.0
	content, err := json.Marshal(user1)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile("practise1.json", content, 0644)
	if err != nil {
		panic(err)
	}

	//Readfile
	datatype, err := ioutil.ReadFile("practise1.json")
	fmt.Println(string(datatype))

	var user2 User
	json.Unmarshal(content, &user2)
	fmt.Println(string(user2.Name))

}
