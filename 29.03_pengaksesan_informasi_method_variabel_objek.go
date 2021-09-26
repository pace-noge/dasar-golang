package main

import (
	"fmt"
	"reflect"
)

type student struct {
	Name  string
	Grade int
}

func (s *student) SetName(name string) {
	s.Name = name
}

func main() {
	var s1 = &student{Name: "John Wick", Grade: 2}
	fmt.Println("name :", s1.Name)

	var reflectValue = reflect.ValueOf(s1)
	var method = reflectValue.MethodByName("SetName")
	method.Call([]reflect.Value{
		reflect.ValueOf("wick"),
	})

	fmt.Println("name :", s1.Name)
}
