//go:build goroutine

package main

import (
	"fmt"
	"time"
)

func sayHello() {
	fmt.Println("Hello, World!")
}
func main() {
	go sayHello()
	fmt.Println("Main thread")
	time.Sleep(1 * time.Second)
}
