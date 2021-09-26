package main

import "fmt"

type person struct {
	name    string
	age     int
	hobbies []string
}

func main() {
	var p1 = struct {
		name string
		age  int
	}{age: 22, name: "Wick"}
	var p2 = struct {
		name string
		age  int
	}{"ethan", 22}

	fmt.Println(p1)
	fmt.Println(p2)
}
