package getHome

import (
	"imageresizerservice/app/ctx/appContext"
	"imageresizerservice/app/home/homeRoutes"
	"imageresizerservice/app/projects/projectRoutes"
	"imageresizerservice/app/ui/page"
	"imageresizerservice/app/users/userAccount/userAccountRoutes"
	"imageresizerservice/library/static"
	"net/http"
)

func Router(mux *http.ServeMux, ac *appContext.AppCtx) {
	mux.HandleFunc(homeRoutes.HomePage, Respond(ac))
}

type Data struct {
	ProjectsPageHref    string
	UserAccountPageHref string
}

func Respond(ac *appContext.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := Data{
			ProjectsPageHref:    projectRoutes.ToProjectListPage(),
			UserAccountPageHref: userAccountRoutes.UserAccountPage,
		}

		page.Respond(static.GetSiblingPath("page.html"), data)(w, r)
	}
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, homeRoutes.HomePage, http.StatusSeeOther)
}
