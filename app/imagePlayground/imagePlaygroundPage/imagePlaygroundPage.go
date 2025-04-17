package imagePlaygroundPage

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/home/homeRoutes"
	"imageresizerservice/app/imagePlayground/imagePlaygroundRoutes"
	"imageresizerservice/app/projects/project"
	"imageresizerservice/app/projects/projectRoutes"
	"imageresizerservice/app/ui/page"
	"imageresizerservice/library/static"
	"net/http"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	mux.HandleFunc(imagePlaygroundRoutes.ImagePlaygroundPage, Respond(ac))
}

type Data struct {
	BackURL          string
	Projects         []project.Project
	CreateProjectURL string
}

func Respond(ac *appCtx.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		data := Data{
			BackURL:          homeRoutes.HomePage,
			Projects:         []project.Project{},
			CreateProjectURL: projectRoutes.ToCreateProject(),
		}

		page.Respond(static.GetSiblingPath("page.html"), data)(w, r)
	}
}

func Redirect(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, homeRoutes.HomePage, http.StatusSeeOther)
}
