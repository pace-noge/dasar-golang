package main

import (
    "fmt"
    "regexp"
)

func main() {
    var text = "banana burger soup"
    var regex, _ = regexp.Compile(text)

    var str = regex.ReplaceAllStringFunc(text, func(each string) string {
        if each == "burger" {
            return "potato"
        }
        return each
    })

    fmt.Println(str)
}
