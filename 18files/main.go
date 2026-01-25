package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Files in golang")

	fmt.Println("Creating a text file in this directory and reading it.")
	fmt.Println("Files are created using: os package (os.Create('where to create the file'))")

	content := "This text came from a golang program"

	file, err := os.Create("./demoTextFile.txt")
	checkNilErr(err)

	length, err := file.WriteString(content)
	checkNilErr(err)

	fmt.Println("Length of the file is :", length)

	defer file.Close()	/// Important

	fmt.Println("Files are read using: os.ReadFile() package")
	readFile("./demoTextFile.txt")
}

func checkNilErr(err error){
	if(err != nil){
		panic(err)
	}
}

func readFile(filename string){
	databyte, err := os.ReadFile(filename)
	checkNilErr(err)

	var stringText = string(databyte)
	fmt.Println("Text file contains: \n", stringText)
}