package main

import (
	"fmt"
)

func main() {
	var s1 = library.student{"ethan", 21}
	fmt.Println("name", s1.name)
	fmt.Println("grade", s1.grade)
}
