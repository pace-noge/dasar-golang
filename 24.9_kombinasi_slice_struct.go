package main

import "fmt"

type Person struct {
    name string
    age int
}

func main() {
    var allStudent = []Person{
        {name: "Wick", age: 23},
        {name: "Ethan", age: 23},
        {name: "Bourne", age: 22},
    }

    for _, student := range allStudent {
        fmt.Println(student.name, "age is", student.age)
    }
}

