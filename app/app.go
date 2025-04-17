package app

import (
	"imageresizerservice/app/api"
	"imageresizerservice/app/apiDocs"
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/home"
	"imageresizerservice/app/home/homePage"
	"imageresizerservice/library/sessionID"
	"imageresizerservice/library/traceID"

	"imageresizerservice/app/projects"
	"imageresizerservice/app/users"
	"imageresizerservice/app/users/auth"
	"imageresizerservice/app/users/login/sendLink/sendLinkPage"
	"imageresizerservice/library/static"
	"net/http"
)

// Handler is the main handler for the application.
func Handler() http.Handler {
	ac := appCtx.New()

	mux := http.NewServeMux()

	router(mux, &ac)

	handler := traceID.WithTraceIDHeader(sessionID.WithSessionIDCookie(mux))

	return handler
}

// router is the router for the application.
func router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	muxLoggedIn := newMuxLoggedIn(ac)
	muxLoggedOut := newMuxLoggedOut(ac)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")

		err := static.ServeStaticAssets(w, r)
		if err == nil {
			return
		}
		if auth.IsLoggedIn(ac, r) {
			muxLoggedIn.ServeHTTP(w, r)
			return
		}
		muxLoggedOut.ServeHTTP(w, r)
	})
	mux.Handle("/", handler)
}

// newMuxLoggedIn is the mux for the logged in user.
func newMuxLoggedIn(ac *appCtx.AppCtx) *http.ServeMux {
	mux := http.NewServeMux()
	users.Router(mux, ac)
	home.Router(mux, ac)
	projects.Router(mux, ac)
	apiDocs.Router(mux, ac)
	api.Router(mux, ac)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		homePage.Redirect(w, r)
	})
	return mux
}

// newMuxLoggedOut is the mux for the logged out user.
func newMuxLoggedOut(ac *appCtx.AppCtx) *http.ServeMux {
	mux := http.NewServeMux()
	users.RouterLoggedOut(mux, ac)
	api.Router(mux, ac)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sendLinkPage.Redirect(w, r)
	})
	return mux
}
