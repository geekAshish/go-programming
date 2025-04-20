package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter your num:")
	input, err := reader.ReadString('\n')

	fmt.Println("It's a gread num", input)

	numRating, err := strconv.ParseFloat(strings.TrimSpace(input), 64)
	if err != nil {
		panic(err)
	}

	fmt.Println("Number rating:", numRating)
}

