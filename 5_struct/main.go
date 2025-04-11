package main

import (
	"fmt"
	"time"
)

type order struct {
	id        string
	amount    float32
	status    string
	createdAt time.Time // Nanoseconds precision
}

func main() {
	myOrder := order {
		id: "1",
		amount: 23.33,
		status: "pending",
	}

	myOrder.createdAt = time.Now()

	fmt.Println(myOrder, myOrder.amount)
}
