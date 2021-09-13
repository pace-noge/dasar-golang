package main

import "fmt"
import "strings"


type Student struct {
    name string
    grade int
}


func (s Student) sayHello() {
    fmt.Println("halo", s.name)
}

func (s Student) getNameAt(i int) string {
    return strings.Split(s.name, " ")[i-1]
}

func main() {
    var s1 = Student{"John Wick", 21}
    s1.sayHello()

    var name = s1.getNameAt(2)
    fmt.Println("last name:", name)
}
