package main

import (
    "fmt"
    "time"
)

func main() {
    fmt.Println("Start")
    <-time.After(4 * time.Second)
    fmt.Println("expired")
}
