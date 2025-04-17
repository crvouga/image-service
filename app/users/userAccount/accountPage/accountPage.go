package accountPage

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/ctx/reqCtx"
	"imageresizerservice/app/home/homeRoutes"
	"imageresizerservice/app/ui/breadcrumbs"
	"imageresizerservice/app/ui/page"
	"imageresizerservice/app/ui/pageHeader"
	"imageresizerservice/app/users/logout/logoutRoutes"
	"imageresizerservice/app/users/userAccount"
	"imageresizerservice/app/users/userAccount/userAccountRoutes"
	"imageresizerservice/app/users/userSession"
	"imageresizerservice/library/static"
	"net/http"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	mux.HandleFunc(userAccountRoutes.UserAccountPage, Respond(ac))
}

type Data struct {
	UserSession *userSession.UserSession
	UserAccount *userAccount.UserAccount
	Breadcrumbs []breadcrumbs.Breadcrumb
	PageHeader  pageHeader.PageHeader
}

func Respond(ac *appCtx.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rc := reqCtx.FromHttpRequest(ac, r)

		data := Data{
			UserSession: rc.UserSession,
			UserAccount: rc.UserAccount,
			Breadcrumbs: []breadcrumbs.Breadcrumb{
				{Label: "Home", Href: homeRoutes.HomePage},
				{Label: "Account"},
			},
			PageHeader: pageHeader.PageHeader{
				Title: "Account",
				Actions: []pageHeader.Action{
					{
						Label: "Logout",
						URL:   logoutRoutes.Logout,
					},
				},
			},
		}

		page.Respond(data, static.GetSiblingPath("accountPage.html"))(w, r)
	}
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, homeRoutes.HomePage, http.StatusSeeOther)
}
