package main

import "fmt"
import "reflect"

func main() {
    var number = 32
    var reflectValue = reflect.ValueOf(number)

    fmt.Println("Variable type :", reflectValue.Type())

    if reflectValue.Kind() == reflect.Int {
        fmt.Println("nilai variable: ", reflectValue.Int())
    }
}
