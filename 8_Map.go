// go:build map

package main

import "fmt"

func main() {
	sehirNufus := map[string]int{
		"Istanbul": 15000000,
		"Ankara":   5000000,
	}
	fmt.Println("Istanbul'un nüfusu: ", sehirNufus["Istanbul"])

	sehirNufus["Istanbul"] = 16000000
	fmt.Println("Istanbul'un nüfusu: ", sehirNufus["Istanbul"])

	delete(sehirNufus, "Ankara")
	fmt.Println("Ankara'nın nüfusu: ", sehirNufus["Ankara"])

	testDb := map[string]struct {
		username string
		mail     string
		age      int
	}{
		"user1": {username: "user1", mail: "user1@example.com", age: 20},
		"user2": {username: "user2", mail: "user2@example.com", age: 25},
	}

	fmt.Println("user1'in maili: ", testDb["user1"].mail)

}
