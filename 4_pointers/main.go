package main

import "fmt"

func main() {
	num := 5

	changeValue(&num)

	fmt.Println(num)
}

func changeValue(a *int) {
	*a = 10

	fmt.Println(a)
}
