package main

import (
	"fmt"
	"time"
)


func someClose() {
	c := make(chan int)

	go func() {
		sum := <-c
		fmt.Println(sum)
	}()

	c <- 111;
}

func bufferedChannels() {
	c := make(chan int, 5)

	go func ()  {
		for i := range 10 {
			fmt.Println("send data", i);

			c <- i;
		}
	}();
	
	time.Sleep(time.Second * 3);
	fmt.Println(<-c)
	// fmt.Println(<-c)
	// fmt.Println(<-c)
}

// Unbuffered channels 
func main() {
	c := make(chan int)

	go func() {
		sum := 0

		for i := range 10000 {
			sum += i
		}

		c <- sum
	}()

	go func() {
		sum := 0

		for i := range 10000 {
			sum += i + 10001
		}

		c <- sum
	}()

	sum1 := <-c
	sum2 := <-c

	fmt.Println(sum1 + sum2)


	// 2
	// someClose()

	// 3
	bufferedChannels()

	// if you want to pass data use channels
	// if you don't want to pass data use sync.WaitGroup
}
