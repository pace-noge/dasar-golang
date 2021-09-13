package main

import "fmt"

type Person struct {
    name string
    age int
}

func main() {
    var s1 = struct{
        Person
        grade int
    }{}

    s1.Person = Person{"wick", 21}
    s1.grade = 2

    fmt.Println("name\t:", s1.Person.name)
    fmt.Println("name\t:", s1.Person.age)
    fmt.Println("name\t:", s1.grade)
}
