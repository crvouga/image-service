package app

import (
	"imageService/app/admin"
	"imageService/app/api"
	"imageService/app/apiDocs"
	"imageService/app/ctx/appCtx"
	"imageService/app/ctx/reqCtx"
	"imageService/app/home"
	"imageService/app/home/homePage"
	"imageService/app/projects"
	"imageService/app/ui/pages"
	"imageService/app/users"
	"imageService/app/users/auth"
	"imageService/app/users/login/sendLink"
	"imageService/library/sessionID"
	"imageService/library/static"
	"imageService/library/traceID"
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
		rc := reqCtx.FromHttpRequest(ac, r)
		rc.Logger.Info("request received", "path", r.URL.Path)

		w.Header().Set("Cache-Control", "no-cache, no-store, must-revalidate")
		w.Header().Set("Pragma", "no-cache")
		w.Header().Set("Expires", "0")

		if err := static.ServeStaticAssets(w, r); err == nil {
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
	pages.Router(mux)
	admin.Router(mux, ac)
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
	pages.Router(mux)
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		sendLink.Redirect(w, r)
	})
	return mux
}
