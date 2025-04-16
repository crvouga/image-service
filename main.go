package main

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/ctx/sessionID"
	"imageresizerservice/app/dashboard"
	"imageresizerservice/app/dashboard/dashboardPage"
	"imageresizerservice/app/users"
	"imageresizerservice/app/users/auth"
	"imageresizerservice/app/users/login/sendLink/sendLinkPage"
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
	loggedInMux := RouterLoggedIn(appCtx)
	loggedOutMux := RouterLoggedOut(appCtx)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")

		err := static.ServeStaticAssets(w, r)
		if err == nil {
			return
		}
		if auth.IsLoggedIn(appCtx, r) {
			loggedInMux.ServeHTTP(w, r)
			return
		}
		loggedOutMux.ServeHTTP(w, r)
	})
	mux.Handle("/", handler)
}

func RouterLoggedIn(appCtx *appCtx.AppCtx) *http.ServeMux {
	mux := http.NewServeMux()
	users.Router(mux, appCtx)
	dashboard.Router(mux, appCtx)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		dashboardPage.Redirect(w, r)
	})

	return mux
}

func RouterLoggedOut(appCtx *appCtx.AppCtx) *http.ServeMux {
	mux := http.NewServeMux()
	users.RouterLoggedOut(mux, appCtx)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sendLinkPage.Redirect(w, r)
	})
	return mux
}
