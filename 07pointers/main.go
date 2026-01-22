package main

import "fmt"

func main() {
	fmt.Println("Pointers in golang")

	one := 1;

	var ptr *int
	fmt.Println("Value of uninitialized pointer (ptr) is: ", ptr); // nil

	var onePointer *int = &one;
	fmt.Println(onePointer); // address of the one variable
	fmt.Printf("Type of onePointer is: %T", *onePointer) // value inside the pointer location

	*onePointer = *onePointer + 1;
	fmt.Println("Value of onePointer after adding 1 is: ", *onePointer, "and value of one is: ", one);
}