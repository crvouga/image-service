package homePage

import (
	"imageresizerservice/app/admin/adminRoutes"
	"imageresizerservice/app/apiDocs/apiDocsRoutes"
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/ctx/reqCtx"
	"imageresizerservice/app/home/homeRoutes"
	"imageresizerservice/app/projects/projectRoutes"
	"imageresizerservice/app/ui/errorPage"
	"imageresizerservice/app/ui/page"
	"imageresizerservice/app/ui/pageHeader"
	"imageresizerservice/app/users/userAccount"
	"imageresizerservice/app/users/userAccount/userAccountRoutes"
	"imageresizerservice/app/users/userAccount/userRole"
	"imageresizerservice/library/static"
	"net/http"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	mux.HandleFunc(homeRoutes.HomePage, Respond(ac))
}

type Data struct {
	ProjectsURL   string
	AccountURL    string
	ApiDocsURL    string
	ClaimAdminURL string
	NoAdmins      bool
	UserAccount   userAccount.UserAccount
	AdminURL      string
	PageHeader    pageHeader.PageHeader
}

func Respond(ac *appCtx.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rc := reqCtx.FromHttpRequest(ac, r)

		admins, err := ac.UserAccountDB.GetByRole(userRole.Admin)

		if err != nil {
			errorPage.New(err).Redirect(w, r)
			return
		}

		data := Data{
			PageHeader: pageHeader.PageHeader{
				Title: "Home",
			},
			ProjectsURL:   projectRoutes.ToListProjects(),
			AccountURL:    userAccountRoutes.UserAccountPage,
			ApiDocsURL:    apiDocsRoutes.ApiDocsPage,
			NoAdmins:      len(admins) == 0,
			ClaimAdminURL: adminRoutes.ClaimAdmin,
			UserAccount:   rc.UserAccount.EnsureComputed(),
			AdminURL:      adminRoutes.AdminPage,
		}

		page.Respond(data, static.GetSiblingPath("homePage.html"))(w, r)
	}
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, homeRoutes.HomePage, http.StatusSeeOther)
}
