package adminPage

import (
	"imageService/app/admin/adminRoutes"
	"imageService/app/ctx/appCtx"
	"imageService/app/home/homeRoutes"
	"imageService/app/ui/breadcrumbs"
	"imageService/app/ui/page"
	"imageService/app/ui/pageHeader"
	"imageService/library/static"
	"net/http"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	mux.HandleFunc(adminRoutes.AdminPage, Respond(ac))
}

type Data struct {
	Breadcrumbs breadcrumbs.Breadcrumbs
	PageHeader  pageHeader.PageHeader
}

func Respond(ac *appCtx.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		data := Data{
			Breadcrumbs: []breadcrumbs.Breadcrumb{
				{Label: "Home", Href: homeRoutes.HomePage},
				{Label: "Admin"},
			},
			PageHeader: pageHeader.PageHeader{
				Title: "Admin",
			},
		}

		page.Respond(data, static.GetSiblingPath("adminPage.html"))(w, r)
	}
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, homeRoutes.HomePage, http.StatusSeeOther)
}
