package main

import "fmt"


func printValues[T int | string](items []T) {
	for _, item := range items {
		fmt.Println(item)
	}
}

func printValuesComparable[T comparable, V string](items []T, name V) {
	for _, item := range items {
		fmt.Println(item, name)
	}
}

type stack[T any] struct {
	elements []T
}

func main() {
	printValues([]int{2, 3, 4})

	myStack := stack[string] {
		elements: []string{"go", "c"},
	}

	fmt.Println(myStack)
}