package logoutPage

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/ctx/reqCtx"
	"imageresizerservice/app/ui/page"
	"imageresizerservice/app/users/logout/logoutRoutes"
	"imageresizerservice/app/users/userSession"
	"imageresizerservice/library/static"
	"net/http"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	mux.HandleFunc(logoutRoutes.LogoutPage, Respond(ac))
}

type Data struct {
	UserSession  *userSession.UserSession
	LogoutAction string
}

func Respond(ac *appCtx.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := reqCtx.FromHttpRequest(ac, r)

		data := Data{
			UserSession:  req.UserSession,
			LogoutAction: logoutRoutes.LogoutAction,
		}

		page.Respond(static.GetSiblingPath("logoutPage.html"), data)(w, r)
	}
}
