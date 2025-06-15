package main

import (
	"fmt"
	"sync"
)

func main() {

	myCh := make(chan int, 2)
	wg := &sync.WaitGroup{}


	wg.Add(2)

	// read ONLY
	go func (ch <-chan int, wg *sync.WaitGroup)  {
		val, isChannelOpen := <-ch
		fmt.Println(isChannelOpen, val)
	}(myCh, wg)

	// send ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 9
		close(ch)
		ch <- 8

		wg.Done()
	}(myCh, wg)

	wg.Wait()
}