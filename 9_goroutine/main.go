package main

import (
	"fmt"
	"time"
)

// for multithreding

func task(num int) {
	fmt.Println("doing task: ", num)
}

func main() {
	for i := 0; i < 10; i++ {
		// go task(i)

		// good practice to take value as arg
		go func(num int) {
			fmt.Println(num)
		} (i)
	}


	time.Sleep(2)
}