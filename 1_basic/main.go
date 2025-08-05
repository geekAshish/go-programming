package main

import (
	"bufio"
	"fmt"
	"maps"
	"os"
	"slices"
	"time"
)

func main() {
	// string, int, flot, bool, bytes
	// int -> int8, int16, int32, int64
	// uint -> uint8, uint16, uint32, uint64, uintptr
	// float -> float32, float64
	// complex -> complex64, complex128
	// Array, slice, map, struct, channel, function, interface, pointer
	fmt.Println(7/3)
	fmt.Println(7.0/3.0)
	// println("Hello World") // println is not recommended, It might get removed from GO, use fmt.Println

	var favNumber float32 = 6.00;
	favNumber = 55
	var favNumberInfer  = 6; // type infer
	favNumberShortHand := 55; // short hand

	// constant
	const username string = "Ashish"

	// const grouping
	const (
		host = "localhost"
		port = 3000
	)
	fmt.Println(favNumber, favNumberInfer, favNumberShortHand, host, port)


	// if : can use variable with condition
	// go does not have ternary
	if age := 122; age > 18 {
		fmt.Println("ashish")
	}

	// comma ok syntax
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your name:")

	name, _ := reader.ReadString('\n')
	fmt.Println("Hello", name)

	// switch_function()
	// forLoop()
	// array_function()
	// slices_function()

	map_function()

}

// for -> only construct in go for looping
func forLoop() {
	// while loop
	i := 0;
	for i < 3 {
		fmt.Println(i)
		i++
	}

	// infinite loop

	// for  {
	// 	fmt.Println("Ashish")
	// }


	//for loop
	for i := 0; i < 10; i++ {
		// break
		// continue
		fmt.Println(i)
	}



	// new feature 1.22 range
	for i := range 5 {
		fmt.Println(i)
	}

}

func switch_function()  {
	// multiple switch
	switch time.Now().Weekday() {
		case time.Saturday, time.Sunday:
			fmt.Println("it's weekend")
		default:
			fmt.Println("it's workday")
	}

	// fallthrough


	// type switch
	whoAmI := func (i interface{})  {
		// switch t := i.(type) {
		switch i.(type) {
		case int:
			fmt.Println("Integer")
		case string:
			fmt.Println("string")
		}
	}

	whoAmI(6)
}

func array_function()  {
	var arr [6]int

	nums := [4]int{1, 2, 3}
	nums2d := [4][4]int{{1, 2, 3}, {1, 2, 3, 4}}

	arr[0] = 2


	// lenght
	fmt.Println(len(arr), nums, nums2d)
}

func slices_function() {
	// uninitialized slice value is nil
	// var num []int;


	// 2. if don't want nil value
	// capacity -> maximum pumbers of elements can fit
	var nums = make([]int, 2, 10)
	nums2 := []int{}

	nums[0] = 1
	nums[1] = 4
	// nums[3] = 4// out of range error, can't add value to index 3, because length is 2
	nums = append(nums, 1) // but use append to add values
	nums = append(nums, 222)
	nums2 = append(nums2, 222)
	fmt.Println(nums, cap(nums), nums2)


	// copy function
	// copy(destination, source)


	// slice operator
	var nums3 = []int {1,32, 45, 6}
	fmt.Println(nums3[0:2])
	
	// slices
	fmt.Println(slices.Equal(nums, nums3))


}

func map_function()  {
	m := make(map[string]string)
	m2 := map[string]string{"game": "socker"}


	m["language"] = "go lang"
	m["price"] = "123"

	// IMP: if key does not exists in the map then it returns zero value
	fmt.Println(m["language"], len(m), m2)

	delete(m, "price")
	clear(m)


	//
	value, ok := m["game"]

	fmt.Println(value)
	if ok {
		fmt.Println("It's OK")
	} else {
		fmt.Println("It's not OK")
	}


	//
	fmt.Println(maps.Equal(m, m2))
}
