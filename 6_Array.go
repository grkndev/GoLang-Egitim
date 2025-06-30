// go:build array

package main

import "fmt"

func main() {
	// var arrayName [size]type
	var squares [5]int

	var numbers [5]int
	fmt.Println("Boş array: ", numbers) // [0 0 0 0 0]

	// Değer atama
	numbers[0] = 10
	numbers[1] = 20
	fmt.Println("Doldurulmuş array: ", numbers) // [10 20 0 0 0]

	for i := 0; i < len(numbers); i++ {
		squares[i] = (i + 1) * (i + 1)
	}
	fmt.Println("Kareler: ", squares) // [1 4 9 16 25]

	var grades []string
	scores := []int{85, 65, 95, 45}

	for _, score := range scores {
		if score >= 90 {
			grades = append(grades, "A")
		} else if score >= 70 {
			grades = append(grades, "B")
		} else {
			grades = append(grades, "F")
		}
	}
	fmt.Println("Notlar: ", grades) // [B F A F]

	exercise()

}

func exercise() {
	// [3, 6, 9, 12, 15] dizindeki çift sayıları yazdıran kodu yazınız.
	numbers := []int{3, 6, 9, 12, 15}
	for _, number := range numbers {
		if number%2 == 0 {
			fmt.Println(number)
		}
	}
}
