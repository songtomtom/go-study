package main

import (
	"fmt"
	"go-study/packages/numbers"
	"go-study/packages/strings"
	"go-study/packages/strings/greeting"
	str "strings"
)

func main() {
	fmt.Println(numbers.IsPrime(19))
	fmt.Println(greeting.WelcomeText)
	fmt.Println(strings.Reverse("callicoder"))
	fmt.Println(str.Count("Go is Awesome. I love Go", "Go"))
}
