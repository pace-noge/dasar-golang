package main

import (
    "fmt"
    "strconv"
)

func main() {
    var str = "124"
    var num, err = strconv.Atoi(str)

    if err == nil {
        fmt.Println(num)
    }

    var numm = 124
    var strr = strconv.Itoa(numm)

    fmt.Println(strr)

    var strrr = "124"
    var nummm, errr = strconv.ParseInt(strrr, 10, 64)

    if errr == nil {
        fmt.Println(nummm)
    }
}
