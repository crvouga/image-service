package main

import (
	"log"
	"net/http"

	"imageresizerservice/deps"
	"imageresizerservice/email/sendEmail"
	"imageresizerservice/static"
	"imageresizerservice/users"
	"imageresizerservice/users/loginEmailLink/loginLink/loginLinkDb"
)

func main() {
	d := deps.Deps{
		SendEmail:   &sendEmail.FakeSendEmail{},
		LoginLinkDb: &loginLinkDb.ImplHashMap{},
	}

	mux := http.NewServeMux()

	mux.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("User-agent: *\nAllow: /"))
	})

	users.Router(mux, &d)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := static.ServeStaticAssets(w, r)
		if err != nil {
			http.Redirect(w, r, "/login-with-email-link/login-page", http.StatusSeeOther)
		}
	})

	port := "8080"

	log.Printf("Server live here http://localhost:%s/ \n", port)

	http.ListenAndServe(":8080", mux)

}
