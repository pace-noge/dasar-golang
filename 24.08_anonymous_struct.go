package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var s1 = struct {
		person
		grade int
	}{}

	s1.person = person{"wick", 21}
	s1.grade = 2

	fmt.Println("name\t:", s1.person.name)
	fmt.Println("age\t:", s1.age)
	fmt.Println("grade\t:", s1.grade)
}
