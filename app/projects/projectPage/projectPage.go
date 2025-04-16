package projectPage

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/dashboard/dashboardRoutes"
	"imageresizerservice/app/projects/project"
	"imageresizerservice/app/projects/project/projectID"
	"imageresizerservice/app/projects/projectRoutes"
	"imageresizerservice/app/ui/page"
	"imageresizerservice/library/static"
	"net/http"
	"net/url"
)

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	mux.HandleFunc(dashboardRoutes.DashboardPage, Respond(appCtx))
}

type Data struct {
	DashboardPage     string
	Project           *project.Project
	EditProjectPage   string
	DeleteProjectPage string
}

func Respond(appCtx *appCtx.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		projectIDMaybe := r.URL.Query().Get("projectID")

		if projectIDMaybe == "" {
			http.Error(w, "Project ID is required", http.StatusBadRequest)
			return
		}

		projectIDInst, err := projectID.New(projectIDMaybe)

		if err != nil {
			http.Error(w, "Invalid project ID", http.StatusBadRequest)
			return
		}

		project, err := appCtx.ProjectDB.GetByID(projectIDInst)

		if err != nil {
			http.Error(w, "Project not found", http.StatusNotFound)
			return
		}

		data := Data{
			DashboardPage:     dashboardRoutes.DashboardPage,
			Project:           project,
			EditProjectPage:   projectRoutes.ProjectEdit,
			DeleteProjectPage: projectRoutes.ProjectDelete,
		}

		page.Respond(static.GetSiblingPath("projectPage.html"), data)(w, r)
	}
}

func Redirect(w http.ResponseWriter, r *http.Request, projectID string) {
	u, _ := url.Parse(projectRoutes.ProjectPage)
	q := u.Query()
	q.Set("projectID", projectID)
	u.RawQuery = q.Encode()
	http.Redirect(w, r, u.String(), http.StatusSeeOther)
}
