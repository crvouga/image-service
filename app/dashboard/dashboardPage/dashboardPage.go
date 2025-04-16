package dashboardPage

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/dashboard/dashboardRoutes"
	"imageresizerservice/app/ui/page"
	"imageresizerservice/library/static"
	"net/http"
)

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	mux.HandleFunc(dashboardRoutes.DashboardPage, Respond(appCtx))
}

func Respond(appCtx *appCtx.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		page.Respond(static.GetSiblingPath("dashboardPage.html"), nil)(w, r)
	}
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, dashboardRoutes.DashboardPage, http.StatusSeeOther)
}
