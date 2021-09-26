package main

import "fmt"

type student struct {
	person struct {
		name string
		age  int
	}
	grade   int
	hobbies []string
}

func main() {
	var s1 = student{}
	s1.person.name = "bilal"
	s1.person.age = 3
	s1.grade = 1
	s1.hobbies = []string{"eat", "cry"}

	fmt.Println(s1.person.name, s1.grade, s1.hobbies)
}
