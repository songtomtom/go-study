package main

import "fmt"

func main() {
	var b byte = 'a'
	var r rune = '♥'
	fmt.Printf("%c = %d and %c =%U\n", b, b, r, r)
}
