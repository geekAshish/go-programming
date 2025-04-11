package main

import "fmt"

// go doesn't have enumerate type
// we can implement using const

// type OrderStatus int

// const (
// 	Received OrderStatus = iota
// 	Confirmed
// 	Prepared
// 	Delivered
// )

type OrderStatus string

const (
	Received OrderStatus = "Received"
	Confirmed = "Confirmed"
	Prepared = "Prepared"
	Delivered = "Delivered"
)

func change0rderStatus(status OrderStatus) {
	fmt.Println("changing order status to" , status)
}

func main() {
	change0rderStatus(Received)
}