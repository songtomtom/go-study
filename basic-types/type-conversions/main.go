package main

import "fmt"

func main() {
	var a int64 = 4
	var b int = int(a)
	var c float64 = 6.4

	var result = float64(b) + c
	fmt.Println(result)

	var i int = 65
	var ui uint = uint(i)
	var f float64 = float64(i)

	fmt.Println(i, ui, f)
}
