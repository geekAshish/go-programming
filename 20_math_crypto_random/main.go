package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

// "math/rand"


func main() {
	// using math
	// rand.Seed(time.Now().UnixNano());
	// fmt.Println(rand.Intn(5) + 1)
	
	// using crypto
	myRandomNum, _ :=  rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(myRandomNum)
}
