package dashboardPage

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/ctx/reqCtx"
	"imageresizerservice/app/dashboard/dashboardRoutes"
	"imageresizerservice/app/projects/projectRoutes"
	"imageresizerservice/app/ui/page"
	"imageresizerservice/app/users/logout/logoutRoutes"
	"imageresizerservice/app/users/userSession"
	"imageresizerservice/library/static"
	"net/http"
)

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	mux.HandleFunc(dashboardRoutes.DashboardPage, Respond(appCtx))
}

type Data struct {
	UserSession       *userSession.UserSession
	LogoutPage        string
	CreateProjectPage string
}

func Respond(appCtx *appCtx.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := reqCtx.FromHttpRequest(appCtx, r)

		data := Data{
			UserSession:       req.UserSession,
			LogoutPage:        logoutRoutes.ToLogoutPage(),
			CreateProjectPage: projectRoutes.ToProjectCreate(),
		}

		page.Respond(static.GetSiblingPath("dashboardPage.html"), data)(w, r)
	}
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, dashboardRoutes.DashboardPage, http.StatusSeeOther)
}
