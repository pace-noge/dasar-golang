package main

import "fmt"


type Student struct {
    name string
    grade int
}

func main() {
    var s1 = Student{}

    s1.name = "John"
    s1.grade = 2

    var s2 = Student{"ethan", 2}

    var s3 = Student{name: "jasno"}

    fmt.Println("Student 1\t:", s1.name)
    fmt.Println("Student 2\t:", s2.name)
    fmt.Println("Stduent 3\t:", s3.name)
}
