package auth

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/ctx/reqCtx"
	"imageresizerservice/app/error/errorPage"
	"net/http"
)

func WithMustBeLoggedIn(appCtx *appCtx.AppCtx) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqCtx := reqCtx.FromHttpRequest(appCtx, r)

			userSession, err := appCtx.UserSessionDb.GetBySessionID(reqCtx.SessionID)

			if err != nil {
				errorPage.Redirect(w, r, err.Error())
				return
			}

			if userSession == nil {
				errorPage.Redirect(w, r, "You must be logged in to access this page")
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func WithMustBeLoggedOut(appCtx *appCtx.AppCtx, redirectTo string, redirectErrorTo string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			reqCtx := reqCtx.FromHttpRequest(appCtx, r)

			userSession, err := appCtx.UserSessionDb.GetBySessionID(reqCtx.SessionID)

			if err != nil {
				errorPage.Redirect(w, r, err.Error())
				return
			}

			if userSession != nil {
				errorPage.Redirect(w, r, "You must be logged out to access this page")
				return
			}

			next.ServeHTTP(w, r)

		})
	}
}
