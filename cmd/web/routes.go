package main

import (
	"net/http"

	"github.com/lorezi/golang-websocket/internal/handlers"

	"github.com/bmizerany/pat"
)

func routes() http.Handler {
	mux := pat.New()

	mux.Get("/", http.HandlerFunc(handlers.Home))
	mux.Get("/ws", http.HandlerFunc(handlers.WSEndPoint))

	return mux
}
