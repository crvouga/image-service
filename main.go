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
)

func main() {
	db := sqlite.New()

	defer db.Close()

	appCtx := appCtx.New(db)

	mux := http.NewServeMux()

	Router(mux, &appCtx)

	handler := reqCtx.WithSessionID(mux)

	addr := ":8080"

	log.Printf("Server live here http://localhost%s/ \n", addr)

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
