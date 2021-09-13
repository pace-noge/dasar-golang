package main

import "fmt"


type Person struct {
    name string
    age int
}

func  main() {
    var student struct {
        Person
        grade int
    }

    student.Person = Person{"wick", 21}
    student.grade = 2

    fmt.Println(student)
}
