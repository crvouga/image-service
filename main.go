package main

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/ctx/reqCtx"
	"imageresizerservice/app/users"
	"imageresizerservice/app/users/loginWithEmailLink/routes"
	"imageresizerservice/library/sqlite"
	"imageresizerservice/library/static"
	"log"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

func main() {
	db := sqlite.New()

	defer db.Close()

	addr := ":8080"

	baseUrl := "http://localhost" + addr

	appCtx := appCtx.New(db, baseUrl)

	mux := http.NewServeMux()

	Router(mux, &appCtx)

	log.Printf("Server live here %s/ \n", baseUrl)

	handler := reqCtx.WithSessionID(mux)

	http.ListenAndServe(addr, handler)
}

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {

	mux.HandleFunc("/robots.txt", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("User-agent: *\nAllow: /"))
	})

	users.Router(mux, appCtx)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := static.ServeStaticAssets(w, r)
		if err != nil {
			http.Redirect(w, r, routes.SendLinkPage, http.StatusSeeOther)
		}
	})

}
