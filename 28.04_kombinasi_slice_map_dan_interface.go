package main

import "fmt"

func main() {
	var person = []map[string]interface{}{
		{"name": "wick", "age": 23},
		{"name": "ethan", "age": 23},
		{"name": "bourne", "age": 22},
	}

	for _, each := range person {
		fmt.Println(each["name"], "age is", each["age"])
	}

	var fruits = []interface{}{
		map[string]interface{}{"name": "strawberry", "total": 10},
		[]string{"mango", "pineaple", "papaya"},
		"orange",
		99,
	}

	for _, each := range fruits {
		fmt.Println(each)
	}

}
