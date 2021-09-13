package main

import "fmt"


type Student struct {
    name string
    grade int
}


func main() {
    var s1 = Student{name: "Wick", grade: 2}
    var s2 *Student = &s1
    fmt.Println("Student 1, name\t:", s1.name)
    fmt.Println("Student 2 name\t:", s2.name)

    s2.name = "ethan"
    fmt.Println("student 1 name\t:", s1.name)
    fmt.Println("student 2 name\t:", s2.name)
}
