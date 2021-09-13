package main

import "fmt"

type Student struct {
    name string
    grade int
}

func main() {
    var s1 Student
    s1.name = "john wick"
    s1.grade = 2


    fmt.Println("name\t:", s1.name)
    fmt.Println("grade\t:", s1.grade)
}
