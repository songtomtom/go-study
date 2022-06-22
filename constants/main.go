package main

import "fmt"

func main() {
	const favoriteLanguage = "Python"
	const sunRisesInTheEast = true

	const country, code = "India", 91

	const (
		id     string  = "E101"
		salary float64 = 50000.0
	)

	const typedInt int = 1234
	const typedString string = "Hi"
	fmt.Println(favoriteLanguage, sunRisesInTheEast, country, code, id, salary, typedInt, typedString)

}
