package main

import (
	"log"
	"net/http"
)

func main() {

	r := routes()

	log.Println("Starting web server on port 8080")

	err := http.ListenAndServe(":8080", r)
	if err != nil {
		log.Panicf("Could not start the server 🥵 %s", err)
	}

}
