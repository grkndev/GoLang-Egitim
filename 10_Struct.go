// go:build struct

package main

import (
	"fmt"
	"time"
)

type User struct {
	ID    int
	Name  string
	Email string
}
type Order struct {
	ID        int
	Name      string
	Price     float64
	OrderedBy User
	OrderedAt time.Time
}

func main() {

	user := User{ID: 1, Name: "GrknDev", Email: "info@grkn.dev"}
	fmt.Println("User: ", user)

	k := User{2, "GrknDev_2", "info@grkn.dev2"}

	var testUser User
	testUser.Name = "GrknDev_3"
	testUser.Email = "info@grkn.dev3"

	fmt.Println("User: ", k.Name)

	new_order := Order{ID: 1, Name: "Order_1", Price: 100.0, OrderedBy: user, OrderedAt: time.Now()}
	fmt.Println("New Order: ", new_order)
}
