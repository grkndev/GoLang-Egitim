// go:build conditions

package main

import "fmt"

func main() {
	point := 65
	excuse := true

	if point >= 50 {
		fmt.Println("Geçtiniz.")
	} else if excuse {
		fmt.Println("İzniniz var.")
	} else {
		fmt.Println("Kaldınız.")
	}
	fmt.Println("--------------------------------")
	if sayi := 10; sayi%2 == 0 {
		fmt.Println("Çift sayı.")
	} else {
		fmt.Println("Tek sayı.")
	}
}
