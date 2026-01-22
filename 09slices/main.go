package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Slices in golang")

	/* 	-> slices
		-> make()
		-> append()
		-> sort()*/

	var studentsList = []string{"John","Tom","Billy"};
	fmt.Printf("Type of studentsList is %T\n",studentsList); // []string

	studentsList = append(studentsList, "Neo","Dave")
	fmt.Println("Adding two more students: ",studentsList)

	studentsList = append(studentsList[:3]) // [inclusive:exclusive]
	fmt.Println("After slicing [:3] :" ,studentsList);

	Marks := make([]int,3)

	Marks[0]=89
	Marks[1]=77
	Marks[2]=85
	// Marks[3]=100 /// throws index out of bound run time error (if used this method to append a new entry)

	Marks = append(Marks, 92,96)
	fmt.Println("Marks: ", Marks)

	fmt.Println(sort.IntsAreSorted(Marks))
	sort.Ints(Marks)
	fmt.Println("After sorting:", Marks)

	/// IMPORTANT - removing a particular item from a slice
	grades := []string{"A","B","C","D"}
	var index = 2
	grades = append(grades[:index],grades[index+1:]...)
	fmt.Println(grades)
}