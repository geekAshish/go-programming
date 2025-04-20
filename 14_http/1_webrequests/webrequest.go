package webrequests

import (
	"fmt"
	"io"
	"net/http"
)

const url string = "https://google.com"

func Webrequests() {
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Response Status Code: %T\n", response)

	defer response.Body.Close() // caller's responsibility to close the connection

	databytes, err := io.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	content := string(databytes)
	fmt.Printf("Response Content: %s\n", content)
}
