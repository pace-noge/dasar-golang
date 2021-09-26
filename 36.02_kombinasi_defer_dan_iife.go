package main

import "fmt"

func main() {
	number := 3

	if number == 3 {
		fmt.Println("halo 1")
		defer fmt.Println("Halo 3")
	}

	fmt.Println("Halo 2")

	number5 := 5

	if number5 == 5 {
		fmt.Println("halo 3")
		func() {
			defer fmt.Println("Halo 5")
		}()
	}

	fmt.Println("halo 4")

}
