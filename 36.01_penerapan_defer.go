package main

import "fmt"

func orderSomeFood(menu string) {
	defer fmt.Println("Terima kasih, silahkan tunggu")
	if menu == "Pizza" {
		fmt.Println("Pilihan Tepat")
		fmt.Println("Pizza di tempat kami yang paling enak", "\n")
		return
	}
}

func main() {
	//	defer fmt.Println("halo")
	fmt.Println("selamat datang")
	orderSomeFood("Pizza")
	orderSomeFood("Burger")

}
