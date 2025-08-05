package main

import (
	"fmt"
	"sync"
)

// while doing mutilthreading, to prevent race condition we use mutex
// lock the variable to the multiple thread don't access it at the same time

type post struct {
	views int
	mu sync.Mutex
}

func (p *post) int(wg *sync.WaitGroup) {
	
	defer func() {
		wg.Done()
		p.mu.Unlock()
	}()

	p.mu.Lock()
	p.views += 1
	// p.mu.Unlock() // put it inside defer so run after func execution end
}

func main() {
	var wg sync.WaitGroup
	post := post{views: 0}

	for i := 0; i < 100; i++ {
		wg.Add(1)
	  go post.int(&wg)
	}
	
	wg.Wait()

	fmt.Println(post.views)
	
}
