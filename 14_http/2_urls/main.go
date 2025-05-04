package main

import (
	"fmt"
	"net/url"
)

const myUrl string = "https://brave.dev:3000/learn?course=go&id=22"

func main() {
	// getting values from the URLs
	result, err := url.Parse(myUrl)

	if err != nil {
		panic("")
	}

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)

	qparams := result.Query()

	fmt.Println(qparams["course"])

	for _, val := range qparams {
		fmt.Println("param is : ", val)
	}


	// constructing an URL from chunks
	partsOfUrl := &url.URL{
		Scheme: "https",
		Host: "lco.dev",
		Path: "/coding",
		RawPath: "user=ashish",
	}

	anotherUrl := partsOfUrl.String()
	fmt.Println(anotherUrl)
}