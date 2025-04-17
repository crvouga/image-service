package apiDocsPage

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/ctx/reqCtx"

	"imageresizerservice/app/home/homeRoutes"
	"imageresizerservice/app/projects/project"
	"imageresizerservice/app/projects/projectRoutes"
	"imageresizerservice/app/ui/errorPage"
	"imageresizerservice/app/ui/page"
	"imageresizerservice/library/static"
	"net/http"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	mux.HandleFunc("/api-docs", Respond(ac))
}

type Data struct {
	HomeURL          string
	Projects         []*project.Project
	CreateProjectURL string
}

func Respond(ac *appCtx.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		rc := reqCtx.FromHttpRequest(ac, r)

		projects, err := ac.ProjectDB.GetByCreatedByUserID(rc.UserSession.UserID)

		if err != nil {
			errorPage.New(err).Redirect(w, r)
			return
		}

		data := Data{
			HomeURL:          homeRoutes.HomePage,
			Projects:         projects,
			CreateProjectURL: projectRoutes.ToCreateProject(),
		}

		page.Respond(data, static.GetSiblingPath("apiDocsPage.html"), "./app/api/apiImageResizer.html")(w, r)
	}
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, homeRoutes.HomePage, http.StatusSeeOther)
}
