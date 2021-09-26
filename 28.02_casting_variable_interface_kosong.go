package main

import (
	"fmt"
	"strings"
)

func main() {
	var secret interface{}

	secret = 2
	var number = secret.(int) * 10
	fmt.Println(secret, "multiplied by 10 is", number)

	secret = []string{"apple", "mango", "banana"}
	// secret casted into its original type, i this case []strings
	var fruits = strings.Join(secret.([]string), ", ")
	fmt.Println(fruits, "is my favorite fruits")

	favFruits := []string{"apple", "mango", "durian"}
	fmt.Println(strings.Join(favFruits, ", "), "is my vaforite fruits")
}
