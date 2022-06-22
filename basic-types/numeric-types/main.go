package main

import "fmt"

func main() {
	var i8 = 97
	var i = 1200
	var ui = 500
	var hex = 0xFF
	var oct = 034
	var f32 float32 = 4.5
	var f = 9.12

	fmt.Printf("%d, %d, %d, %#x, %#o, %f, %f\n", i8, i, ui, hex, oct, f32, f)
}
