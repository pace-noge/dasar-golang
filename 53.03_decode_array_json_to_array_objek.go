package main

import (
    "fmt"
    "encoding/json"
)


type User struct {
    FullName string `json:Name`
    Age int
}


func main() {
    var jsonString = `[
        {"Name": "john wick", "Age": 27},
        {"Name": "ethan hunt", "Age": 32}
    ]`

    var data []User

    var err = json.Unmarshal([]byte(jsonString), &data)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println("User 1:", data[0].FullName)
    fmt.Println("User 2:", data[1].FullName)
}
