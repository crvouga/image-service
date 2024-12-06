package main

import (
	"log"
	"net/http"

	"imageresizerservice/deps"
	"imageresizerservice/email/sendEmail"
	"imageresizerservice/static"
	"imageresizerservice/users/loginEmailLink"
	"imageresizerservice/users/loginEmailLink/routes"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("User-agent: *\nAllow: /"))
	})

	d := deps.Deps{
		SendEmail: &sendEmail.FakeSendEmail{},
	}

	loginEmailLink.Router(mux, &d)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Println(r.URL.Path)
		err := static.ServeStaticAssets(w, r)
		if err != nil {
			http.Redirect(w, r, routes.Prefix, http.StatusSeeOther)
		}
	})

	port := "8080"

	log.Printf("Server live here http://localhost:%s/ \n", port)

	http.ListenAndServe(":8080", mux)

}
