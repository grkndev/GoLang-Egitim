// go:build pointer

/*
	Pointer: Tanımlanmış bir değişkenin bellekteki adresini gösteren değerdir.
	Neden?
		- int türünde parametre alan fonksyion çalıştırdığımızda değer bellekte kopyalanır ve kopya değer üzerinden işlem yapılır.
		Lakin Pointer ile fonksyion çalıştırdığımızda bu sefer parametre olarak tam sayı değil referans değerini göndeririz yani aslında değişkenin DEĞERİNİ değil, değişkenin KENDİSİNİ göndeririz.
		Böylelikle sistem bellekte kopyalama yapmaz ve referans değer üzerinden işlem yapılır.
*/

package main

import "fmt"

func main() {
	var test_value int = 10
	test_value_p := &test_value

	fmt.Println("test_value: ", test_value)                              // 10
	fmt.Println("test_value_p: ", test_value_p)                          // test_value'un adresi
	fmt.Println("test_value_p'nin işaret ettiği değer: ", *test_value_p) // test_value'un değeri | 10
	fmt.Println("--------------------------------")

	artir(&test_value)

	fmt.Println("--------------------------------")
	fmt.Println("test_value: ", test_value)                              // 11
	fmt.Println("test_value_p'nin işaret ettiği değer: ", *test_value_p) // test_value'un değeri | 11
	fmt.Println("test_value_p: ", test_value_p)                          // test_value'un adresi
}

func artir(x *int) {
	fmt.Println("x'in adresi: ", x) // X zaten bellekteki değeri temsil eder.
	// fmt.Println("x'in adresinin adresi: ", &x)
	fmt.Println("x'in işaret ettiği değer: ", *x)
	*x = *x + 1
}
