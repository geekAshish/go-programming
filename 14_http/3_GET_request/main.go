package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	PerformGetRequest()
}

func PerformGetRequest() {
	const myUrl = "http://localhost:3000/get"

	response, err := http.Get(myUrl);
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	fmt.Println("status code", response.StatusCode)
	fmt.Println("content length", response.ContentLength)

	var responseString strings.Builder
	content, _ := io.ReadAll(response.Body)

	byteCount, _ := responseString.Write(content)

	// one way
	fmt.Println(string(content))

	// another way, with more options
	fmt.Println("bytecount is :", byteCount)
	fmt.Println(responseString.String())
}