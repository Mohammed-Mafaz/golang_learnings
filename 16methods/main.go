package main

import "fmt"

func main() {
	fmt.Println("Methods in golang")

	// a function within a struct is called method

	structure := User{"John Doe","john@go.dev",25}
	structure.GetAge()
	structure.updateMail()

	fmt.Println("[main()] --- Email:",structure.Email)	// doesn't change

	// var a = 10;
	// var b = 20;
	// swap(a,b);
	// fmt.Printf("a: %d, b:%d \n", a,b)

}

type User struct{
	Name string
	Email string
	Age int
}

func (u User) GetAge(){
	fmt.Println("Age of the user is :", u.Age)
}

// func swap (a int,b int){
// 	var temp = a;
// 	a = b;
// 	b = temp;
// }

func (u User) updateMail(){
	u.Email = "newMail@go.dev"
	fmt.Println("[updateMail()] --- new Email:",u.Email)
}
