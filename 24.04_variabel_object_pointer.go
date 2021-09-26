package main

import "fmt"

type student struct {
	name  string
	grade int
}

func main() {
	var s1 = student{name: "Wick", grade: 2}

	var s2 *student = &s1

	fmt.Println("Student 1 name:", s1.name)
	fmt.Println("Student 2 name:", s2.name)

	s2.name = "ethan"
	fmt.Println("Student 1 name:", s1.name)
	fmt.Println("Student 2 name:", s2.name)
}
