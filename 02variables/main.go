package main

import "fmt"


func main() {
	var name string = "John Doe"
	fmt.Println("Name: ", name);
	fmt.Printf("Variable type: %T \n" , name);
	var isLoggedIn bool = true;
	fmt.Println("isLoggedIn: " , isLoggedIn);
	fmt.Printf("Variable type: %T \n", isLoggedIn);

	var small uint8= 255;
	fmt.Println(small);
	fmt.Printf("Variable type: %T \n", small);

   var smallFloat float32= 255.345678990; // 5 decimal values
	fmt.Println(smallFloat);
	fmt.Printf("Variable type: %T \n", smallFloat);


	var integer int= 987654321987654321;
	fmt.Println(integer);
	fmt.Printf("Variable type: %T \n", integer);

	var notAssignedInteger int;
	fmt.Println(notAssignedInteger); /// Value will be 0;


	/// implicit type
	
	var greeting = "Greetings fellow Humans!";
	fmt.Println(greeting);

	// can't change the type of the variable later.
	// greeting = 10; -- Not allowed


	/// no var style (volarous operater)
	numberOfUsers := 3000.0;
	fmt.Println(numberOfUsers)
	numberOfUsers = 10;
	fmt.Println(numberOfUsers)

	//-- can't use volarous operator Outside the function.

	/// Public and Private variables
	fmt.Println(A_Public_Variable);	
	fmt.Println(a_Private_Variable)
}

var A_Public_Variable string = "I am a Public variable String";	// public variable

var a_Private_Variable string = "I am a Private variable String"; // private variable