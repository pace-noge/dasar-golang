package main

import "fmt"


type Person struct {
    name string `tag1`
    age int     `tag2`
}

func main() {
    var p1 = Person{
        name: "bilal",
        age: 3,
    }
    fmt.Println(p1)
}
