package projectCreate

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/projects/projectRoutes"
	"imageresizerservice/app/ui/page"
	"imageresizerservice/library/static"
	"net/http"
)

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	mux.HandleFunc(projectRoutes.ProjectCreate, Respond(appCtx))
}

type Data struct {
}

func Respond(appCtx *appCtx.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			respondPost(appCtx, w, r)
		} else {
			respondGet(w, r)
		}
	}
}

func respondGet(w http.ResponseWriter, r *http.Request) {
	data := Data{}
	page.Respond(static.GetSiblingPath("projectCreate.html"), data)(w, r)
}

func respondPost(appCtx *appCtx.AppCtx, w http.ResponseWriter, r *http.Request) {
	// Handle form submission
	r.ParseForm()

	// Redirect back to the form or to a success page
	http.Redirect(w, r, projectRoutes.ProjectCreate, http.StatusSeeOther)
}
