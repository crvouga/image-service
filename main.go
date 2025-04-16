package main

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/ctx/sessionID"
	"imageresizerservice/app/dashboard"
	"imageresizerservice/app/users"
	"imageresizerservice/app/users/login/loginRoutes"
	"imageresizerservice/library/static"
	"log"
	"net/http"
)

func main() {
	appCtx := appCtx.New()

	mux := http.NewServeMux()

	Router(mux, &appCtx)

	handler := sessionID.WithSessionIDCookie(mux)

	addr := ":8080"

	log.Printf("Server live here http://localhost%s/ \n", addr)

	http.ListenAndServe(addr, handler)
}

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	users.Router(mux, appCtx)
	dashboard.Router(mux, appCtx)

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := static.ServeStaticAssets(w, r)
		if err != nil {
			http.Redirect(w, r, loginRoutes.SendLinkPage, http.StatusSeeOther)
		}
	})
}
