package main

import (
    "fmt"
    "encoding/json"
)


type User struct {
    Name string `json:FullName`
    Age int
}


func main() {
    var object = []User{{"john wick", 27}, {"ethan hunt", 32}}
    var jsonData, err = json.Marshal(object)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    var jsonString = string(jsonData)
    fmt.Println(jsonString)
}
