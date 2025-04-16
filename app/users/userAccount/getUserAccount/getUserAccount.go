package getUserAccount

import (
	"imageresizerservice/app/ctx/appContext"
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

func Router(mux *http.ServeMux, ac *appContext.AppCtx) {
	mux.HandleFunc(userAccountRoutes.UserAccountPage, Respond(ac))
}

type Data struct {
	UserSession *userSession.UserSession
	LogoutPage  string
	UserAccount *userAccount.UserAccount
	BackURL     string
}

func Respond(ac *appContext.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		reqCtxInst := reqCtx.FromHttpRequest(ac, r)

		data := Data{
			UserSession: reqCtxInst.UserSession,
			UserAccount: reqCtxInst.UserAccount,
			LogoutPage:  logoutRoutes.ToLogoutPage(),
			BackURL:     homeRoutes.HomePage,
		}

		page.Respond(static.GetSiblingPath("page.html"), data)(w, r)
	}
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, homeRoutes.HomePage, http.StatusSeeOther)
}
