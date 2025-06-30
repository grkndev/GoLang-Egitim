// go:build checkpoint1

package main

import "fmt"

func main() {
	size := 850
	maxsize := 1024
	percent := (float32(size) / float32(maxsize)) * 100
	if percent < 50 {
		fmt.Printf("Alan kullanımı düşük: %.2f%%\n", percent)
	} else if percent < 80 {
		fmt.Printf("Alan kullanımı orta: %.2f%%\n", percent)
	} else if percent < 100 {
		fmt.Printf("Alan kullanımı yüksek: %.2f%%\n", percent)
	} else {
		fmt.Printf("Alan kullanımı çok yüksek: %.2f%%\n", percent)
	}
}
