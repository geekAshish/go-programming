package main

import "fmt"

func main() {
	var myNumber int = 20;
	var prtNum *int;
	
	prtNum = &myNumber;
	
	num := 5
	changeValue(&num)

	fmt.Println(num, &myNumber, prtNum, *prtNum + *prtNum)
}

func changeValue(a *int) {
	*a = 10

	fmt.Println(a)
}
