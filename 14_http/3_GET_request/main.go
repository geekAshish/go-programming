package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	// PerformGetRequest()
	PerformPostFormRequest()
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

func PerformPostJsonRequest() {
	const url = ""

	// fake json data
	requestBody := strings.NewReader(`
		{
			"name": "ashish"
		}
	`)

	response, err := http.Post(url, "application/json", requestBody)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}

func PerformPostFormRequest() {
	const myUrl = "http://localhost:3000/postform";

	// form data

	data := url.Values{}

	data.Add("firstname", "ashish");
	data.Add("userid", "47")

	response, err := http.PostForm(myUrl, data)
	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	content, _ := io.ReadAll(response.Body)

	fmt.Println(string(content))
}