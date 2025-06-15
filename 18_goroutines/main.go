package goroutines

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup // usualy these are pointers

func main() {
	greet("hey")
	greet("bro")

	websiteList := []string {"github.com", "google.com", "fb.com"};

	for _, web := range websiteList {
		go getStatusCode(web)
		wg.Add(1)
	}

	wg.Wait()
}

func greet(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}

// wait group
func getStatusCode(endpoint string) {
	wg.Done()
	
	res, err := http.Get(endpoint)

	if err != nil {
		fmt.Println("OOPS in endpoint")
	}

	fmt.Println(res.StatusCode, endpoint)
}