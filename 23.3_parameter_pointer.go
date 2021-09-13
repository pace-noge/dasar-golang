package main

import "fmt"

func main() {
    var number = 4
    fmt.Println("before: ", number)

    change(&number, 10)
    fmt.Println("after\t:", number)
}

func change(original *int, value int) {
    *original = value
}
