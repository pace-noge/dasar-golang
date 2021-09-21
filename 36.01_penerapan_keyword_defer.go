package main

import "fmt"

func orderSomeFood(menu string) {
	defer fmt.Println("Terima kasih, silahkan tunggu")
	if menu == "pizza" {
		fmt.Println("Pilihan Tepat", " ")
		fmt.Println("Pizza ditempat kami paling enak!", "\n")
		return
	}
}

func main() {

	// // this line will be executed last
	// defer fmt.Println("halo")
	// fmt.Println("Selamat Datang")
	orderSomeFood("pizza")
	orderSomeFood("burger")
}
