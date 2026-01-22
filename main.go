package main

import (
	"imageService/app"
	"os"

	"log"
	"net/http"
)

func main() {

	handler := app.Handler()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	addr := ":" + port

	log.Printf("Server live here http://localhost%s/ \n", addr)

	http.ListenAndServe(addr, handler)
}
