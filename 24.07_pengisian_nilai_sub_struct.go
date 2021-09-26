package main

import "fmt"

type person struct {
	name string
	age  int
}

type student struct {
	person
	grade int
}

func main() {
	var p1 = person{name: "Wick", age: 21}
	var s1 = student{person: p1, grade: 2}

	fmt.Println("name\t:", s1.name)
	fmt.Println("age\t:", s1.age)
	fmt.Println("grade\t:", s1.grade)
}
