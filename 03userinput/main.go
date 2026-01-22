package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	
	welcome := "Welcome to user input";
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin);
	fmt.Println("Enter a number");
	
	// comma ok || comma err syntax
	input, _ := reader.ReadString('\n');	// "\n - wrong; expect a character|byte"
	
	fmt.Println("You entered: ", input)
	fmt.Printf("Type of the input value of the variable is: %T", input)


}