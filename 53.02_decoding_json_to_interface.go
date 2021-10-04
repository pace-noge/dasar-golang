package main

import (
    "fmt"
    "encoding/json"
)

func main() {
    var jsonString = `{"Name": "john wick", "Age": 21}`
    var jsonData = []byte(jsonString)

    var data1 map[string]interface{}
    json.Unmarshal(jsonData, &data1)

    fmt.Println("User\t:", data1["Name"])
    fmt.Println("Age\t:", data1["Age"])


    var data2 interface{}
    json.Unmarshal(jsonData, &data2)
    var decodedData = data2.(map[string]interface{})
    fmt.Println("user\t:", decodedData["Name"])
    fmt.Println("age\t:", decodedData["Age"])
    
}
