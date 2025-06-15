package goroutines

import "fmt"

func main() {
	greet("hey")
	greet("bro")
}

func greet(s string) {
	for i := 0; i < 5; i++ {
		fmt.Println(s)
	}
}
