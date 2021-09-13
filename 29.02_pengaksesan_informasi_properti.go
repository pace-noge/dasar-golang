package main

import "fmt"
import "reflect"

type student struct {
    Name string
    Grade int
}

func (s *student) getPropertyInfo() {
    var reflectValue = reflect.ValueOf(s)

    if reflectValue.Kind() == reflect.Ptr {
        reflectValue = reflectValue.Elem()
    }

    var reflectType = reflectValue.Type()

    for i := 0; i < reflectValue.NumField(); i++ {
        fmt.Println("name\t:", reflectType.Field(i).Name)
        fmt.Println("tipe data\t:", reflectType.Field(i).Type)
        fmt.Println("nilai\t:", reflectValue.Field(i).Interface())
        fmt.Println("")
    }
}

func main() {
    var s1 = &student{Name: "Wick", Grade: 2}
    s1.getPropertyInfo()
}
