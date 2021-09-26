package main

import (
	"fmt"
	"reflect"
)

func main() {
	var number = 23
	var reflectValue = reflect.ValueOf(number)

	fmt.Println("tipe variabel:", reflectValue.Type())

	if reflectValue.Kind() == reflect.Int {
		fmt.Println("nilai variable:", reflectValue.Int())
	}

	// akses nilai dengan interface
	fmt.Println("nilai lewat interface", reflectValue.Interface())

	var nilai = reflectValue.Interface().(int)

	fmt.Println(nilai)
}
