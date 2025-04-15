package main

import (
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"

	"imageresizerservice/deps"
	"imageresizerservice/sqlite"
	"imageresizerservice/static"
	"imageresizerservice/users"
	"imageresizerservice/users/loginWithEmailLink/routes"
)

func main() {
	db := sqlite.New()

	defer db.Close()

	addr := ":8080"

	baseUrl := "http://localhost" + addr

	d := deps.New(db, baseUrl)

	mux := http.NewServeMux()

	Router(mux, &d)

	log.Printf("Server live here %s/ \n", baseUrl)

	http.ListenAndServe(addr, mux)
}

func Router(mux *http.ServeMux, d *deps.Deps) {

	mux.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("User-agent: *\nAllow: /"))
	})

	users.Router(mux, d)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := static.ServeStaticAssets(w, r)
		if err != nil {
			http.Redirect(w, r, routes.SendLinkPage, http.StatusSeeOther)
		}
	})
}
