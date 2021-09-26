package main

import "fmt"

type student struct {
	name  string
	grade int
}

func main() {
	var s1 = student{}
	s1.name = "John"
	s1.grade = 2

	var s2 = student{"ethan", 2}

	var s3 = student{name: "jason"}

	fmt.Println("Student 1\t:", s1.name)
	fmt.Println("Student 2\t:", s2.name)
	fmt.Println("Student 3\t:", s3.name)
}
