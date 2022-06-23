package main

import "fmt"

func main() {

	var fileExtensions = map[string]string{
		"Python": ".py",
		"C++":    ".cpp",
		"Java":   ".java",
		"Golang": ".go",
		"Kotlin": ".kt",
	}

	fmt.Println(fileExtensions)
	delete(fileExtensions, "Kotlin")
	delete(fileExtensions, "Javascript")
	fmt.Println(fileExtensions)

}
