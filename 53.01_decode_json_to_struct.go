package main

import (
    "encoding/json"
    "fmt"
)

type User struct {
    FullName string `json:"NAME"`
    Age int
}

func main() {
    var jsonString = `{"Name": "John Wick", "Age": 27}`
    var jsonData = []byte(jsonString)

    var data User

    var err = json.Unmarshal(jsonData, &data)
    if err != nil {
        fmt.Println(err.Error())
        return
    }

    fmt.Println("user\t:", data.FullName)
    fmt.Println("age\t:", data.Age)
}
