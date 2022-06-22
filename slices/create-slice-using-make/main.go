package main

import "fmt"

func main() {
	s := make([]int, 5, 10)
	fmt.Printf("s = %v, len = %d, cap = %d\n", s, len(s), cap(s))

	var s1 = make([]int, 5)
	fmt.Printf("s1 = %v, len = %d, cap = %d\n", s1, len(s1), cap(s1))
}
