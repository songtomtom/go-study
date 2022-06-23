package main

import (
	"fmt"
	"go-study/structs/struct-exported-vs-unexported/model"
)

type Employee struct {
	Id   int
	Name string
}

func main() {

	c := model.Customer{
		Id:   1,
		Name: "Rajeev Singh",
	}

	fmt.Println("Programmer = ", c)

}
