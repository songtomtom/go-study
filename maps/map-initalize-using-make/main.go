package main

import "fmt"

func main() {
	var m = make(map[string]int)
	fmt.Println(m)
	if m == nil {
		fmt.Println("m is nil")
	} else {
		fmt.Println("m is not nil")
	}

	m["one hundred"] = 100
	fmt.Println(m)
}
