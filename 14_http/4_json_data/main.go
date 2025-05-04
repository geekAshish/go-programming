package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name  string `json:"name"`
	Price int `json:"price"`
	Email string `json:"email"`
	Password string `json:"-"`
	Tags  []string `json:"tags,omitempty"`
}

func main() {
	EncodedJson()
}

func EncodedJson() {
	courses := []course{
		{"reactjs", 222, "@mail", "222", []string{"web"}},
		{"go", 322, "@mail", "222", []string{"back-web"}},
		{"rust", 422, "@mail", "222", nil},
	}

	// package this data in JSON Data

	// finalJson, err := json.Marshal(courses)
	finalJson, err := json.MarshalIndent(courses, "", "\t")

	if err != nil {
		panic(err)
	}

	fmt.Printf("%s\n", finalJson)

}

