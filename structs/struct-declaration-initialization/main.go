package main

import "fmt"

type Person struct {
	FirstName string
	LastName  string
	Age       int
}

func main() {
	var p Person
	fmt.Println(p)

	p1 := Person{"Rajeev", "Singh", 26}
	fmt.Println("Person1: ", p1)

	p2 := Person{
		FirstName: "John",
		LastName:  "Snow",
		Age:       45,
	}
	fmt.Println("Person2: ", p2)

	p3 := Person{FirstName: "Robert"}
	fmt.Println("Person3: ", p3)

}
