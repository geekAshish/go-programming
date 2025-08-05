package main

import "fmt"

func main() {
	nums := []int{6, 4, 8}
	sum := 0

	for index, num := range nums {
		sum += num
		fmt.Println(sum, index)
	}

	for range 5 {
		fmt.Println("Hello")
	}


	// with map
	player := map[string]string{"name": "ashish", "class": "older"}

	for key, value := range player {
		fmt.Println(key, value)
	}

	// string, unicode code point rune
	// starting byte of rune
	for i, c := range "ashish" {
		fmt.Println(i, c)
	}
}


