package main

import (
	"fmt"
	"math"
)

func main() {
	var a, b = 4, 5
	var res1 = (a + b) * (a + b) / 2
	a++
	b += 10
	var res2 = a ^ b
	var r = 3.5
	var res3 = math.Pi * r * r

	fmt.Printf("res1: %v, res2: %v, res3: %v\n", res1, res2, res3)
}
