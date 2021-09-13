package main

import "fmt"
import "strings"

func filter(data []string, callback func(string) bool) []string {
    var result []string
    for _, each := range data {
        if filtered := callback(each); filtered {
            result = append(result, each)
        }
    }
    return result
}

func main() {
    var data = []string{"wick", "json", "ethan"}
    var dataContainsO = filter(data, func(each string) bool {
        return strings.Contains(each, "o")
    })

    var dataLength5 = filter(data, func(each string) bool {
        return len(each) == 5
    })

    fmt.Println("Real data\t\t:", data)
    fmt.Println("filter ada huruf \"o\"\t:", dataContainsO)

    fmt.Println("Filter jumlah huruf \"5\"\t:", dataLength5)
}

