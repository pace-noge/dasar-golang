package main

import (
    "fmt"
    "regexp"
)


func main() {
    var text = "banan burger soup"
    var regex, _ = regexp.Compile(`[a-b]+`)

    var str = regex.Split(text, -1)
    fmt.Printf("%#v\n", str)
}
