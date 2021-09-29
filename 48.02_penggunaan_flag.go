package main

import (
    "flag"
    "fmt"
)

func main() {
    var name = flag.String("name", "anonymous", "type your name")
    var age = flag.Int64("age", 25, "Type your age")

    flag.Parse()
    fmt.Printf("name\t: %s\n", *name)
    fmt.Printf("age\t: %d\n", *age)

    // go run 48.02_penggunaa_flag.go -name="john wick" -age=28
}
