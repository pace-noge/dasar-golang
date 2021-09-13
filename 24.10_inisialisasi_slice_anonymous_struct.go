package main

import "fmt"

type Person struct {
    name string
    age int
}

func main() {
    var allStudents = []struct {
        Person
        grade int
    }{
        {Person: Person{"wick", 21}, grade: 2},
        {Person: Person{"ethan", 22}, grade: 3},
        {Person: Person{"bond", 21}, grade: 3},
    }

    for _, stuDent := range allStudents {
        fmt.Println(student)
    }
}
