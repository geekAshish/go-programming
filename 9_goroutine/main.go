package main

import (
	"fmt"
	"sync"
)

// for multithreding

func task(num int, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("doing task: ", num)
}

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go task(i, &wg)

		// good practice to take value as arg
		// go func(num int) {
		// 	fmt.Println(num)
		// } (i)
	}


	wg.Wait()
}