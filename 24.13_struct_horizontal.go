package  main

import "fmt"

type Person struct { name string; age int; hobbies []string }


func main() {
    var p1 = struct { name string; age int }{ age: 22, name: "Wick" }
    fmt.Println(p1)
}
