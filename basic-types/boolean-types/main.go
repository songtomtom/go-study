package main

import "fmt"

func main() {
	var b1 bool = true
	var b2 = false

	var truthy = 3 <= 5
	var falsy = 10 != 10

	var res1 = 10 > 20 && 5 == 5
	var res2 = 2*2 == 4 || 10%3 == 0

	fmt.Println(b1, b2, truthy, falsy, res1, res2)
}
