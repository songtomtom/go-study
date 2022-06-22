package main

import "fmt"

func main() {
	var s = []int{3, 5, 7, 9, 11, 13, 17} // Creates an array, and returns a slice reference to the array
	t := []int{2, 4, 8, 16, 32, 64}
	fmt.Println("s = ", s)
	fmt.Println("t = ", t)
}
