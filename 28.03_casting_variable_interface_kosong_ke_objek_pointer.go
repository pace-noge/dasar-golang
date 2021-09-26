package main

import "fmt"

type person struct {
	name string
	age  int
}

func main() {
	var secret interface{} = &person{name: "john", age: 27}
	var name = secret.(*person).name
	fmt.Println(name)
}
