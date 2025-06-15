package main

import (
	"log"
	"net/http"

	"github.com/geekAshish/mongodb/router"
)

func main() {
	r := router.Router()

	err := http.ListenAndServe(":4000", r)

	if err != nil {
		log.Fatal(err)
	}
}