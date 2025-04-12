package main

import (
	"fmt"
	"sync"
)

// while doing mutilthreading, to prevent race condition we use mutex

type post struct {
	views int
}

func (p *post) int(wg *sync.WaitGroup) {
	defer wg.Done()
	p.views += 1
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
