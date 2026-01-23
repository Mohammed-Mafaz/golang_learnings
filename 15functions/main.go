package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Functions in Golang")
	greet()

	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Type of the variable 'reader' is: %T\n", reader)

	fmt.Println("Let's add two numbers")

	fmt.Print("Enter operand 1: ")
	op1Str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read operand 1")
		return
	}

	op1, err := strconv.ParseInt(strings.TrimSpace(op1Str), 10, 64)
	if err != nil {
		fmt.Println("Invalid number for operand 1")
		return
	}

	fmt.Print("Enter operand 2: ")
	op2Str, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Failed to read operand 2")
		return
	}

	op2, err := strconv.ParseInt(strings.TrimSpace(op2Str), 10, 64)
	if err != nil {
		fmt.Println("Invalid number for operand 2")
		return
	}

	result := add(op1, op2)
	fmt.Printf("%d + %d = %d\n", op1, op2, result)

	
	TotalSum := proAdder(10,20,30,40,50)
	fmt.Println("Total sum from proAdder is:",TotalSum)

	FinalSum, Message := twoResults(1,2,3,4,5)
	fmt.Printf("FinalSum: %d and Message: %s", FinalSum,Message)
}

func greet() {
	fmt.Println("Greetings from Golang")
}

func add(a int64, b int64) int64 {
	return a + b
}

func proAdder(values ...int) int{
	var Total int
	for _,val := range values{
		Total += val
	}
	return Total
}

func twoResults(values ...int) (int,string){
	var Total int
	for _,val := range values{
		Total += val
	}
	return Total, "Total is Computed"
}


