package main

import "fmt"

func main() {
	fmt.Println("Structs in golang")
	println()

	obj := User{"John Doe", "dummy@mail.com", 33}
	fmt.Println(obj)
	fmt.Printf("Struct contains: %+v\n", obj)
	fmt.Printf("Name: %v, email:%v.", obj.Name, obj.Email)

}

type User struct {
	Name   string
	Email  string
	Age    int
}

/* Note: No inheritance, super, parent exists in golang*/