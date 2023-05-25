package main

import (
	"net/http"
)

func main() {
	mux := routes()

	_ = http.ListenAndServe(":8080", mux)
}
