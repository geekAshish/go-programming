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

func DecodedJson() {
	jsonDataFromWeb := []byte(`
	{
		"name": "ashish",
		"price": 22,
		"email": "@mail",
		"password": "33",
		tags: ["aa", "ww"]
	}
	`)

	var courses course

	checkValid := json.Valid(jsonDataFromWeb)

	if checkValid {
		fmt.Println("JSON IS VALID")
		json.Unmarshal(jsonDataFromWeb, &courses)
		fmt.Printf("%#v\n", courses);
	} else {
		fmt.Println("JSON IS NOT VALID")
	}


	// some cases where you just want to add data to key value


	var myOnlineData map[string]interface{}
	json.Unmarshal(jsonDataFromWeb, &myOnlineData)
	fmt.Printf("%#v\n", myOnlineData);

	for key, value := range myOnlineData {
		fmt.Printf("key is %v and value is %v and type is %T\n", key, value, value)
	}

}