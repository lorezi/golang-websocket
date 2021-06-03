package main

import (
	"log"
	"net/http"

	"github.com/lorezi/golang-websocket/internal/wsocket"
)

func main() {

	r := routes()

	log.Println("Starting channel listener")

	go wsocket.ListenToWSChannel()

	log.Println("Starting web server on port 8088 🤝")

	err := http.ListenAndServe(":8088", r)
	if err != nil {
		log.Panicf("Could not start the server 🥵 %s", err)
	}

}
