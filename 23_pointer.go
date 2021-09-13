package main

import "fmt"

func main() {
    var numberA int = 4
    var numberB *int = &numberA

    fmt.Println("numberA (value)\t:", numberA)
    fmt.Println("numberA (address)\t:", &numberA)
    fmt.Println("numberB (value)\t:", *numberB)
    fmt.Println("numberB (address)\t:", numberB)
}

