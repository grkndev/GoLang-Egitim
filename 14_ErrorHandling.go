//go:build errorhandling

package main

import (
	"errors"
	"fmt"
)

func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, errors.New("Bölme işlemi sıfıra bölünemez")
	}
	return a / b, nil
}

type CustomError struct {
	Code    int
	Message string
}

func (e CustomError) Error() string {
	return fmt.Sprintf("Hata kodu: %d, Mesaj: %s", e.Code, e.Message)
}

func fail() error {
	return CustomError{Code: 404, Message: "Sayfa bulunamadı"}
}

func main() {
	result, err := divide(10, 0)
	if err != nil {
		fmt.Println("Hata:", err)
	} else {
		fmt.Println("Sonuç:", result)
	}
}

func myPanic() {
	panic("Panik mesajı")
}

func safeCall() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Kurtarıldı:", r)
		}
	}()
	myPanic()
}
