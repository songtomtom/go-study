package main

import "fmt"

func main() {
	var a = [5]int{2, 4, 6, 8, 10}
	fmt.Println(a)

	b := [5]int{2, 4, 6, 8, 10}
	fmt.Println(b)

	c := [5]int{2}
	fmt.Println(c)

	d := [...]int{3, 5, 7, 9, 11, 13, 17}
	fmt.Println(d)

}
