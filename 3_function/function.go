package main

import "fmt"

func main() {
	sum := add(1, 2)
	fmt.Println(sum)

  // Returning Multiple Values
	v1, v2, v3 := returnMultipleValues()
	fmt.Println(v1, v2, v3)

	// storing function in variable
	anotherFunc := func (a int) int  {
		return a
	}

	// Take Function
	takeFuncParam(anotherFunc)
	
	// Return Function
	fn := returnFuncParam()
	fmt.Println(fn(400))

	
	// Variadic function
	numbersSlice := []int{1,2,3,4,5}
	// addNumbers(1,2,3,4,5)
	addNumbers(numbersSlice...)


	// Clousers 
	counterFunction := counter()

	counterFunction()
	counterFunction()
	counterFunction()
	counterSum := counterFunction()
	fmt.Println(counterSum)

}

// for same type param : a, b int
func add(a int, b int) int {
	return a + b
}

func returnMultipleValues() (string, string, int) {
	return "golang", "typescript", 2
}

func takeFuncParam(fn func(a int) int) {
	fn(2)
}

func returnFuncParam() func(a int ) int {
	return func(a int ) int {
		return a
	}
}

// Variadit Function
// Can take any number of paramertes
func addNumbers(nums ...int) {
	sum := 0

	for _, value := range nums {
		sum += value
	}

	fmt.Println(sum)
	// return sum
}


// Clousers, like JavaScript
func counter () func() int {
	count := 0

	return func () int {
		count += 1
		return count
	}
}