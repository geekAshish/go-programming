package main

import (
	"fmt"
	"time"
)

type customer struct {
	name string
	phone string

}

// composition
type order struct {
	id        string
	amount    float32
	status    string
	customer
	createdAt time.Time // Nanoseconds precision
}

// constructor function
func newOrder(id string, amount float32, status string) *order {
	// initial function goes here...
	myOrder := order {
		id: id,
		amount: amount,
		status: status,
	}

	return &myOrder
}

// receiver type
func (o *order) changeStatus(status string) {
	o.status = status
}
// just to get value no need to use pointers
func (o order) getAmount() float32 {
	return o.amount
}

func main() {
	// if you don't set any field, default value is zero value

	// customer := customer {
	// 	name: "Ashish",
	// 	phone: "3`22333",
	// }

	myOrder := order {
		id: "1",
		amount: 23.33,
		status: "pending",
		// customer: customer,
		customer: customer {
			name: "Ashish",
			phone: "3`22333",
		},
	}

	// myOrder := newOrder("123", 12.21, "pending")

	// If you just want to use struct only one time
	myLanguage := struct {
		name string
		isFast bool
	} {name: "go", isFast: true}

	myOrder.createdAt = time.Now()

	myOrder.changeStatus("confirm")

	fmt.Println(myOrder, myOrder.amount, myOrder.status, myOrder.getAmount(), myLanguage, myLanguage.name, myOrder.customer,  myOrder.customer.name)
	fmt.Printf("%+v", myOrder)
}
