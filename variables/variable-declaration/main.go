package main

import "fmt"

func main() {
	// Declaring variables
	var s string = "Hello"
	var i int = 0
	var f float64 = 45.45
	fmt.Println(s, i, f)

	// Multiple declarations
	var (
		id                  int    = 5
		firstName, lastName string = "tom", "song"
	)
	fmt.Println(id, firstName, lastName)

	// Short variable declartion syntax
	name := "tom song"
	age, salary, isHappy := 36, 50000.0, true

	fmt.Println(name, age, salary, isHappy)

}
