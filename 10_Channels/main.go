package main

import (
	"fmt"
	"time"
)

func processNum(NumChan chan int) {
	for s := range NumChan {
		fmt.Println("Processing Num", s)
		time.Sleep(time.Second)
	}
	// fmt.Println("Processing Num", <-NumChan)
}

func sum(result chan int, num1 int, num2 int)  {
	sumResult := num1 + num2
	result <- sumResult
}

// goroutine synchronizer
func task(done chan bool) {
	defer func() {done <- true}()

	fmt.Println("Processing...")
}


// we add type safty for sending and receiving data
// func emailSender(emailChan <-chan string, done chan<- bool) {
func emailSender(emailChan chan string, done chan bool) {
	defer func(){ done <- true}()

	for email := range emailChan {
		fmt.Println("Sending email to", email)

		time.Sleep(time.Second)
	}
}

func main() {
	messageChan := make(chan int)
	result := make(chan int)

	go processNum(messageChan)
	// channels are blocking
	// for {
	// 	messageChan <- rand.Intn(100)
	// }
	




	go sum(result, 4, 3)
	res := <- result
	fmt.Println(res)




	// done := make(chan bool)
	// go task(done)
	// <- done // block





	// Buffer channel
	emailChan := make(chan string, 100)
	done := make(chan bool)

	go emailSender(emailChan, done)

	for i := 0; i < 5; i++ {
		emailChan <- fmt.Sprintf("%d@email.com", i)
	}

	fmt.Println("done sending...")

	// It is importent to close channel
	close(emailChan)
	<-done






	// Multiple channels
	chan1 := make(chan int)
	chan2 := make(chan string)

	go func() {
		chan1 <- 1
	}()

	go func() {
		chan2 <- "ashish"
	}()

	for i := 0; i < 2; i++ {
		select {
		case chan1Value := <- chan1:
			fmt.Println("recieved data from chan1", chan1Value)
		case chan2Value := <- chan2:
			fmt.Println("recieved data from chan2", chan2Value)
		}
	}
}
