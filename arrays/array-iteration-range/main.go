package main

import "fmt"

func main() {
	daysOfWeek := [7]string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}

	for index, value := range daysOfWeek {
		fmt.Printf("Day %d of week = %s\n", index, value)
	}

	a := [4]float64{3.5, 7.2, 4.8, 9.5}
	sum := float64(0)

	for _, value := range a {
		sum = sum + value
	}

	fmt.Printf("Sum of all the elements in array %v = %f\n", a, sum)

}
