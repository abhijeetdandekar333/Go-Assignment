package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to files in go lang")
	content := "This has to go in a file"

	file, err := os.Create("./myGoFile.txt")
	if err != nil {
		panic(err) //shutdown execution of program and shows the error
	}

	io.WriteString(file, content)

	length, err := io.WriteString(file, content)

	if err != nil {
		panic(err)
	}
	fmt.Println("Length is: ", length)
	defer file.Close()
	readFile("./myGoFile.txt")
}

func readFile(filename string) {

	databyte, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	fmt.Println("Text data inside file is :\n", string(databyte))

}
