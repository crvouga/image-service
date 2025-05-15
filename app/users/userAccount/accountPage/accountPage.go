package accountPage

import (
	"imageService/app/ctx/appCtx"
	"imageService/app/ctx/reqCtx"
	"imageService/app/home/homeRoutes"
	"imageService/app/ui/breadcrumbs"
	"imageService/app/ui/page"
	"imageService/app/ui/pageHeader"
	"imageService/app/users/logout/logoutRoutes"
	"imageService/app/users/userAccount"
	"imageService/app/users/userAccount/userAccountRoutes"
	"imageService/app/users/userSession"
	"imageService/library/static"
	"net/http"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	mux.HandleFunc(userAccountRoutes.UserAccountPage, Respond(ac))
}

type Data struct {
	UserSession *userSession.UserSession
	UserAccount *userAccount.UserAccount
	Breadcrumbs breadcrumbs.Breadcrumbs
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
