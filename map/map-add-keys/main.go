package main

import "fmt"

func main() {
	var tinderMatch = make(map[string]string)

	tinderMatch["Rajeev"] = "Anelina"
	tinderMatch["James"] = "Sophia"
	tinderMatch["David"] = "Emma"

	fmt.Println(tinderMatch)

	tinderMatch["Rajeev"] = "Jennifer"
	fmt.Println(tinderMatch)
}
