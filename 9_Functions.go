// go:build functions

package main

import "fmt"

/*
func isim(parametreler) dönüşTipi { return değer }
*/
func topla(a int, b int) int { return a + b }

/*
Çoklu dönüş değeri (Multiple Return)
func isim(parametreler) (dönüşTipi1, dönüşTipi2, ...) { return değer1, değer2, ... }
*/
func bolVeKalan(x, y int) (int, int) {
	return x / y, x % y
}

/*
Değişken gibi kullanılabilir
f := func(a, b int) int { return a + b }
*/

/*
Fonksyion parametresi olarak
func uygula(x int, islem func(int) int) int {return islem(x)}
*/
func uygula(x int, islem func(int) int) int { return islem(x) }
func main() {
	fmt.Println("Toplama: ", topla(10, 20))

	bolum, kalan := bolVeKalan(10, 3)
	fmt.Println("Bölüm: ", bolum, "Kalan: ", kalan)

	f := func(a, b int) int { return a + b }
	fmt.Println("Fonksiyon: ", f(10, 80))

	fmt.Println("Uygula: ", uygula(10, func(x int) int { return x * x }))
}
