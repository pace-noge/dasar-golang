package main

import "fmt"

type Person struct {
	name string
	age  int
}

type people = Person

func main() {
	var p1 = Person{"wick", 21}
	fmt.Println(p1)

	var p2 = people{"wick", 21}
	fmt.Println(p2)

	type Number = int

	var num Number = 21
	fmt.Println(num)
}
