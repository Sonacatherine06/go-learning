package main

import "fmt"

type Student struct {
	Name  string
	Age   int
	Grade string
}

func main() {
	s := Student{
		Name:  "Sona",
		Age:   21,
		Grade: "A",
	}

	fmt.Println("Name :", s.Name)
	fmt.Println("Age  :", s.Age)
	fmt.Println("Grade:", s.Grade)
}
