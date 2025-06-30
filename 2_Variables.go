//go:build variables

package main

import "fmt"

/* Değişkenler

Tip		| Açıklama						  | Örnek
-------------------------------------------------------------------------------
string 	| Metin 						  | "Merhaba"
int 	| Tam sayı 						  | 123
float 	| Ondalıklı sayı 				  | 3.14
bool	| Doğru veya yanlış 			  | true
array 	| Sabit uzunluktaki veri grubu	  | [1, 2, 3]
slice 	| Değişken uzunluktaki veri grubu | []int{1, 2, 3}
map 	| Anahtar-değer çiftleri		  | map[string]int{"a": 1, "b": 2}
struct 	| Özel veri yapısı 				  | struct { Name string; Age int }
pointer | Bellek adresi 			 	  | &var


go'ya özgü kısa tanımlama yöntemi
--> yas := 25
':=' ile tip belirtmeden değişken tanımlanabilir.

*/

func main() {
	const pi = 3.14
	var name string = "Grkn"
	var age int = 25
	var height float32 = 1.85
	var isStudent bool = true
	var tags = []string{"go", "python", "java"} //Array
	// fmt.Println(tags[0]) //go
	var tags2 = map[string]int{"go": 1, "python": 2, "java": 3} //Map

	var tags3 = struct {
		Name string
		Age  int
	}{"Grkn", 25} //Struct
	/*
		tags3.Name reutrn Grkn
		tags3.Age  return 25
	*/

	var tags4 = &name //Pointer

	fmt.Println(pi, name, age, height, isStudent, tags, tags2, tags3, tags4)
}
