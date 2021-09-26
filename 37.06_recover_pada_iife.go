package main

import "fmt"

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("error occured", r)
		} else {
			fmt.Println("Application running perfectly")
		}
	}()

	panic("Some Error happen")
}
