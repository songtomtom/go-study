package main

import "fmt"

func main() {
	// Type inference
	var name = "tom song"
	fmt.Printf("Variable 'name' is of type %T\n", name)
	// Multiple variable declarations with inferred types
	var firstName, lastName, age, salary = "tom", "song", 36, 5000.0
	fmt.Printf("firstName: %T, lastName: %T, age: %T, salary: %T\n",
		firstName, lastName, age, salary)
}
