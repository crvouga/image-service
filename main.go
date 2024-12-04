package main

import (
	login_page "image-resizer-service/login"
	"log"
	"net/http"
)

func main() {
	mainMux := http.NewServeMux()

	mainMux.Handle("/", login_page.Routes())

	port := "8080"

	log.Printf("Server live here http://localhost:%s/ \n", port)

	http.ListenAndServe(":8080", mainMux)

}
