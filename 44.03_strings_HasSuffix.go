package main

import (
    "fmt"
    "strings"
)

func main() {
    isSuffix1 := strings.HasSuffix("john wick", "ic")
    fmt.Println(isSuffix1)

    var isSuffix2 = strings.HasSuffix("jhon wick", "ck")
    fmt.Println(isSuffix2)
}
