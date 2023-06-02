package main

import (
	"chat_server/handlers"
	"log"
	"net/http"
)

func main() {
	mux := routes()

	log.Println("Started channel listener")

	go handlers.ListenToWsChannel()

	log.Println("Started web server")

	_ = http.ListenAndServe(":8080", mux)
}
