package main

import "fmt"

func main() {
	fmt.Println("Arrays in golang")

	var list [3]string;	// made use using 0-based indexing
	list[0] = "Apple";
	// list[1] = "Orange"; /// space
	list[2] = "Grapes";

	fmt.Println("List contains: ", list)
	fmt.Println("Length of the list is: ", len(list))	// 1-based count

	var subjects = [4]string{"math","geography","aviations","physics"}
	fmt.Println("Subjects: " ,subjects)
}