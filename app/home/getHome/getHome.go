package getHome

import (
	"imageresizerservice/app/apiDocs/apiDocsRoutes"
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/home/homeRoutes"
	"imageresizerservice/app/projects/projectRoutes"
	"imageresizerservice/app/ui/page"
	"imageresizerservice/app/users/userAccount/userAccountRoutes"
	"imageresizerservice/library/static"
	"net/http"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	mux.HandleFunc(homeRoutes.HomePage, Respond(ac))
}

type Data struct {
	ProjectsURL string
	AccountURL  string
	ApiDocsURL  string
}

func Respond(ac *appCtx.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := Data{
			ProjectsURL: projectRoutes.ToListProjects(),
			AccountURL:  userAccountRoutes.UserAccountPage,
			ApiDocsURL:  apiDocsRoutes.ApiDocsPage,
		}

		page.Respond(static.GetSiblingPath("page.html"), data)(w, r)
	}
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, homeRoutes.HomePage, http.StatusSeeOther)
}
