package main

import "fmt"
import "reflect"


type student struct {
    Name string
    Grade int
}

func (s *student) SetName(name string) {
    s.Name = name
}

func main() {
    var s1 = &student{Name: "John", Grade: 2}
    fmt.Println("name\t\t:", s1.Name)

    var reflectValue = reflect.ValueOf(s1)
    var method = reflectValue.MethodByName("SetName")
    method.Call([]reflect.Value{
        reflect.ValueOf("Wick"),
    })

    fmt.Println("nama\t\t:", s1.Name)
}
