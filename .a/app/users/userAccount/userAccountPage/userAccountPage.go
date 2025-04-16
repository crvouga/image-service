package userAccountPage

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/ctx/reqCtx"
	"imageresizerservice/app/home/homeRoutes"
	"imageresizerservice/app/ui/page"
	"imageresizerservice/app/users/logout/logoutRoutes"
	"imageresizerservice/app/users/userAccount"
	"imageresizerservice/app/users/userAccount/userAccountRoutes"
	"imageresizerservice/app/users/userSession"
	"imageresizerservice/library/static"
	"net/http"
)

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	mux.HandleFunc(userAccountRoutes.UserAccountPage, Respond(appCtx))
}

type Data struct {
	UserSession *userSession.UserSession
	LogoutPage  string
	UserAccount *userAccount.UserAccount
	HomePage    string
}

func Respond(appCtx *appCtx.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqCtxInst := reqCtx.FromHttpRequest(appCtx, r)

		data := Data{
			UserSession: reqCtxInst.UserSession,
			UserAccount: reqCtxInst.UserAccount,
			LogoutPage:  logoutRoutes.ToLogoutPage(),
			HomePage:    homeRoutes.HomePage,
		}

		page.Respond(static.GetSiblingPath("userAccountPage.html"), data)(w, r)
	}
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, homeRoutes.HomePage, http.StatusSeeOther)
}
