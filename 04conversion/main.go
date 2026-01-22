package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	fmt.Print("Enter a number in range 1-5 :");
	reader := bufio.NewReader(os.Stdin);
	input,_ := reader.ReadString('\n');
	fmt.Println("You entered: ", input);

	fmt.Println("Adding 1 to your number");
	inputNumber, err := strconv.ParseFloat(strings.TrimSpace(input),64);

	if err == nil {	// no error case
		if(inputNumber >=1 && inputNumber <= 5){
			inputNumber++;
			fmt.Println("New number is :", inputNumber)
		}else{
			fmt.Println("Please enter a number in range 1-5")
		}
		
	}else{
		fmt.Println("Please enter a valid number.")
	}
}