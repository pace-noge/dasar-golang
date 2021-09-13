package main

import "fmt"


type Person struct {
    name string
    age int
}

type Student struct {
    Person
    age int
    grade int
}

func main() {
    var s1 = Student{}

    s1.name = "wick"
    s1.age = 21
    s1.Person.age = 22

    fmt.Println(s1.name)
    fmt.Println(s1.age)
    fmt.Println(s1.Person.age)
}
