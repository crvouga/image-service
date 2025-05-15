package main

import (
	"imageService/app"

	"log"
	"net/http"
)

func main() {

	handler := app.Handler()

	addr := ":8080"

	log.Printf("Server live here http://localhost%s/ \n", addr)

	http.ListenAndServe(addr, handler)
}
