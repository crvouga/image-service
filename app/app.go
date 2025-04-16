package app

import (
	"imageresizerservice/app/api"
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/home"
	"imageresizerservice/app/home/getHome"
	"imageresizerservice/app/imageResizer"
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
	appCtx := appCtx.New()

	mux := http.NewServeMux()

	router(mux, &appCtx)

	handler := traceID.WithTraceIDHeader(sessionID.WithSessionIDCookie(mux))

	return handler
}

// router is the router for the application.
func router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	muxLoggedIn := newMuxLoggedIn(appCtx)
	muxLoggedOut := newMuxLoggedOut(appCtx)
	handler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")

		err := static.ServeStaticAssets(w, r)
		if err == nil {
			return
		}
		if auth.IsLoggedIn(appCtx, r) {
			muxLoggedIn.ServeHTTP(w, r)
			return
		}
		muxLoggedOut.ServeHTTP(w, r)
	})
	mux.Handle("/", handler)
}

// newMuxLoggedIn is the mux for the logged in user.
func newMuxLoggedIn(appCtx *appCtx.AppCtx) *http.ServeMux {
	mux := http.NewServeMux()
	users.Router(mux, appCtx)
	home.Router(mux, appCtx)
	projects.Router(mux, appCtx)
	imageResizer.Router(mux, appCtx)
	api.Router(mux, appCtx)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		getHome.Redirect(w, r)
	})
	return mux
}

// newMuxLoggedOut is the mux for the logged out user.
func newMuxLoggedOut(appCtx *appCtx.AppCtx) *http.ServeMux {
	mux := http.NewServeMux()
	users.RouterLoggedOut(mux, appCtx)
	api.Router(mux, appCtx)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sendLinkPage.Redirect(w, r)
	})
	return mux
}
