// go:build slices

package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	numbers = append(numbers, 4, 5)
	fmt.Println("numbers: ", numbers) // [1 2 3 4 5]

	a := []int{1, 2, 3}
	b := []int{4, 5, 6}
	a = append(a, b...)
	fmt.Println("a: ", a) // [1 2 3 4 5 6]

	orjinal := []int{1, 2, 3}
	kopya := make([]int, len(orjinal))
	copy(kopya, orjinal)
	fmt.Println("orjinal: ", orjinal) // [1 2 3]
	fmt.Println("kopya: ", kopya)     // [1 2 3]
}
