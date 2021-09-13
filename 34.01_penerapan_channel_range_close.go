package main

import "fmt"

func sendMessage(ch chan<- string) {
	for message := range ch {
		fmt.Println(message)
	}
}
