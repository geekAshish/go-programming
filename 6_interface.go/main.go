package main

import "fmt"

type paymenter interface {
	pay(amount float32) float32
	refund(accountId string)
}

type payment struct {
	getway paymenter
}

func (p payment) makePayment(amount float32) {
	p.getway.pay(amount)
}

type razorpay struct{}

func (r razorpay) pay(amount float32) float32 {
	fmt.Println("making payment user razorpay", amount)
	return amount
}
func (r razorpay) refund(accountId string) {
	fmt.Println("payment refund initiated for ", accountId)
}

type strip struct{}

func (s strip) pay(amount float32) float32 {
	fmt.Println("making payment user strip", amount)
	return amount
}

func main() {

	razorpayGatway := razorpay{}
	// stripGatway := strip{}

	newPayment := payment {
		getway : razorpayGatway,
	}

	newPayment.makePayment(100)

}
