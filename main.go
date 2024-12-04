package main

import (
	"image-resizer-service/deps"
	"image-resizer-service/email/send_email"
	"image-resizer-service/login"
	"image-resizer-service/login/login_routes"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("User-agent: *\nAllow: /"))
	})

	d := deps.Deps{
		SendEmail: &send_email.FakeSendEmail{},
	}

	mux.Handle(login_routes.Prefix, login.Router(d))

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		http.Redirect(w, r, login_routes.Prefix, http.StatusSeeOther)
	})

	port := "8080"

	log.Printf("Server live here http://localhost:%s/ \n", port)

	http.ListenAndServe(":8080", mux)

}
