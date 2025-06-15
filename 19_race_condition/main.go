package main

import (
	"fmt"
	"sync"
)

// To check your programe having a race condition
// go run --race .

func main() {
	wg := &sync.WaitGroup{}
	mut := &sync.RWMutex{} // use where read/ write happen

	score := []int{0}

	wg.Add(4)
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("One R")
		mut.Lock()
		score = append(score, 1)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Two R")
		mut.Lock()
		score = append(score, 2)
		mut.Unlock()
		wg.Done()
	}(wg, mut)

	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Three R")
		mut.Lock()
		score = append(score, 3)
		mut.Unlock()
		wg.Done()
	}(wg, mut)
	
	go func(wg *sync.WaitGroup, m *sync.RWMutex) {
		fmt.Println("Read R")
		mut.RLock()
		fmt.Println(score)
		mut.RUnlock()
		wg.Done()
	}(wg, mut)

	wg.Wait()
	fmt.Println(score)
}