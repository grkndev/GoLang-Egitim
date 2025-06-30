// go:build loop

package main

import "fmt"

func main() {
	// Klasik for döngüsü
	for i := 0; i < 5; i++ {
		fmt.Println("(For) Step: ", i)
	}
	fmt.Println("--------------------------------")
	// while benzeri döngü
	x := 0
	for x < 5 {
		fmt.Println("(While) Step: ", x)
		x++
	}
	fmt.Println("--------------------------------")
	// Sonsuz döngü
	for {
		fmt.Println("(Infinite) Sonsuz döngü")
		break // Sonsuz döngüyü sonlandırır
	}
	fmt.Println("--------------------------------")
	// break ve continue
	for i := 0; i < 10; i++ {
		if i == 3 {
			continue
		}
		fmt.Println("(Break & Continue) Step: ", i)
	}
	fmt.Println("--------------------------------")
	// for range
	liste := []string{"elma", "armut", "kiraz"}
	for i, meyve := range liste {
		fmt.Println("(For Range) Index: ", i, "Value: ", meyve)
	}
}
