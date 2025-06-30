//go:build interfaces

package main

import (
	"fmt"
	"math"
)

// 1. TEMEL INTERFACE TANIMI
// Interface, methodların imzalarını tanımlar. (implementasyon değil)

type Shape interface {
	Area() float64
	Perimeter() float64
}

// 2. STRUCT'LAR VE INTERFACE IMPLEMENTASYONU
// Go'da explicit implements kelimesi yoktur
// Bir struct, interface'in tüm methodlarını tanımlıyorsa otomatik o interface'i implemente eder.

type Rectangle struct {
	Width, Height float64
}

// Rectangle için Shape interface methodlarını tanımlıyoruz
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Circle struct {
	Radius float64
}

// Circle için Shape interface methodlarını tanımlıyoruz
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 3. INTERFACE'LERİN KULLANIMI
// Interface parametresi alan fonksiyon
func PrintShapeInfo(s Shape) {
	fmt.Printf("Alan: %f, Çevre: %f\n", s.Area(), s.Perimeter())
}

//  4. EMPTY INTERFACE
//     Herhangi bir tipte değer alabilir. (Java'da Object gibi)
func PrintAnything(value interface{}) {
	fmt.Printf("Değer: %v, Tip: %T\n", value, value)
}

// 5. TYPE ASSERTION
// interface{} tipindeki değeri gerçek tipine dönüştürme
func ProcressValue(value interface{}) {
	// Güvenli type assertion
	if str, ok := value.(string); ok {
		fmt.Printf("String: %s (uzunluk: %d)\n", str, len(str))
	} else if num, ok := value.(int); ok {
		fmt.Printf("Integer: %d\n", num)
	} else {
		fmt.Printf("Bilinmeyen tip: %T\n", value)
	}
}

// 6. TYPE SWITCH
// Birden fazla tip kontrolü
func TypeSwitch(value interface{}) {
	switch v := value.(type) {
	case string:
		fmt.Printf("String: %s (uzunluk: %d)\n", v, len(v))
	case int:
		fmt.Printf("Integer: %d\n", v)
	case float64:
		fmt.Printf("Float: %f\n", v)
	case bool:
		fmt.Printf("Boolean: %t\n", v)
	default:
		fmt.Printf("Bilinmeyen tip: %T\n", v)
	}
}

// 7. INTERFACE COMPOSITION (Interface Embedding)
type Writer interface {
	Write([]byte) (int, error)
}

type Reader interface {
	Read([]byte) (int, error)
}

type ReadWriter interface {
	Reader
	Writer
}

func main() {
	// Rectangle ve Shape oluşturma
	rect := Rectangle{Width: 10, Height: 5}
	circle := Circle{Radius: 3}

	fmt.Println("=== Interface Kullanımı ===")
	PrintShapeInfo(rect)
	PrintShapeInfo(circle)

	fmt.Println("=== Empty Interface ===")
	PrintAnything("Merhaba")
	PrintAnything(123)
	PrintAnything(true)
	PrintAnything(rect)

	fmt.Println("=== Type Assertion ===")
	ProcressValue("Hello")
	ProcressValue(123)
	ProcressValue(3.14)

	fmt.Println("=== Type Switch ===")
	values := []interface{}{"hello", 42, 3.14, true, rect}
	for _, val := range values {
		TypeSwitch(val)
	}

	fmt.Println("=== Interface Slice ===")
	shapes := []Shape{rect, circle, Rectangle{2, 8}, Circle{5}}
	for i, shape := range shapes {
		fmt.Printf("Şekil %d: Alan: %f, Çevre: %f\n", i, shape.Area(), shape.Perimeter())
		PrintShapeInfo(shape)
	}

}
