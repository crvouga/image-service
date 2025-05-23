package apiDocsPage

import (
	"imageService/app/api"
	"imageService/app/apiDocs/apiDocsRoutes"
	"imageService/app/ctx/appCtx"
	"imageService/app/ctx/reqCtx"
	"imageService/app/home/homeRoutes"
	"imageService/app/ui/breadcrumbs"
	"imageService/app/ui/mainMenu"
	"imageService/app/ui/page"
	"imageService/app/ui/pageHeader"
	"imageService/library/static"
	"net/http"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	mux.HandleFunc("/api-docs", Respond(ac))
}

const (
	PageTitle = "HTTP API Docs"
)

func Respond(ac *appCtx.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		endpoint := r.URL.Query().Get("endpoint")
		rc := reqCtx.FromHttpRequest(ac, r)
		switch endpoint {
		case api.EndpointApiImageResize:
			type Data struct {
				PageHeader      pageHeader.PageHeader
				Breadcrumbs     []breadcrumbs.Breadcrumb
				ProjectIDSelect ProjectIDSelect
				Endpoint        string
				ExampleImageURL string
			}

			data := Data{
				Endpoint:        api.EndpointApiImageResize,
				ExampleImageURL: rc.BaseURL + "dog.jpeg",
				ProjectIDSelect: GetProjectIDSelect(ac, r),
				PageHeader: pageHeader.PageHeader{
					Title:   api.EndpointApiImageResize,
					Actions: []pageHeader.Action{},
				},
				Breadcrumbs: []breadcrumbs.Breadcrumb{
					{Label: "Home", Href: homeRoutes.HomePage},
					{Label: PageTitle, Href: apiDocsRoutes.ApiDocsPage},
					{Label: api.EndpointApiImageResize},
				},
			}

			page.Respond(data, static.GetSiblingPath("apiImageResizer.html"), static.GetSiblingPath("projectIDSelect.html"))(w, r)
			return
		default:

			type Data struct {
				PageHeader  pageHeader.PageHeader
				Breadcrumbs []breadcrumbs.Breadcrumb
				MainMenu    mainMenu.MainMenu
			}

			data := Data{
				PageHeader: pageHeader.PageHeader{
					Title:   PageTitle,
					Actions: []pageHeader.Action{},
				},
				Breadcrumbs: []breadcrumbs.Breadcrumb{
					{Label: "Home", Href: homeRoutes.HomePage},
					{Label: PageTitle},
				},
				MainMenu: mainMenu.MainMenu{
					Items: []mainMenu.MainMenuItem{
						{
							Label:       api.EndpointApiImageResize,
							URL:         apiDocsRoutes.ToApiDocsPage(api.EndpointApiImageResize),
							Description: "Resize an image",
						},
					},
				},
			}

			page.Respond(data, static.GetSiblingPath("apiDocsPage.html"))(w, r)
		}

	}
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, homeRoutes.HomePage, http.StatusSeeOther)
}
