package main

import "fmt"


type Person struct {
    name string
    age int
}


type Student struct {
    grade int
    Person
}

func main() {
    var s1 = Student{}
    s1.name = "Wick"
    s1.age = 21
    s1.grade = 2

    fmt.Println("Name\t:", s1.name)
    fmt.Println("age\t:", s1.age)
    fmt.Println("age\t:", s1.Person.age)
    fmt.Println("grade\t:", s1.grade)
}

