package projectPage

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/ctx/reqCtx"
	"imageresizerservice/app/projects/project"
	"imageresizerservice/app/projects/project/projectID"
	"imageresizerservice/app/projects/projectRoutes"
	"imageresizerservice/app/ui/page"
	"imageresizerservice/library/static"
	"net/http"
)

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	mux.HandleFunc(projectRoutes.ProjectPage, Respond(appCtx))
}

type Data struct {
	BackURL   string
	Project   *project.Project
	EditURL   string
	DeleteURL string
}

func Respond(appCtx *appCtx.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := reqCtx.FromHttpRequest(appCtx, r)
		logger := req.Logger

		logger.Info("projectPage", "projectID", r.URL.Query().Get("projectID"))

		projectIDMaybe := r.URL.Query().Get("projectID")

		logger.Info("projectIDMaybe", "projectIDMaybe", projectIDMaybe)

		if projectIDMaybe == "" {
			logger.Error("missing project ID", "error", "Project ID is required")
			http.Error(w, "Project ID is required", http.StatusBadRequest)
			return
		}

		logger.Info("projectIDMaybe", "projectIDMaybe", projectIDMaybe)

		projectIDNew, err := projectID.New(projectIDMaybe)

		if err != nil {
			logger.Error("invalid project ID", "error", err)
			http.Error(w, "Invalid project ID", http.StatusBadRequest)
			return
		}

		uow, err := appCtx.UowFactory.Begin()
		if err != nil {
			logger.Error("database access failed", "error", err)
			http.Error(w, "Failed to access database", http.StatusInternalServerError)
			return
		}
		defer uow.Rollback()

		project, err := appCtx.ProjectDB.GetByID(projectIDNew)

		if err != nil {
			logger.Error("project not found", "projectID", projectIDMaybe, "error", err)
			http.Error(w, "Project not found", http.StatusNotFound)
			return
		}

		if project == nil {
			logger.Error("project not found", "projectID", projectIDMaybe)
			http.Error(w, "Project not found", http.StatusNotFound)
			return
		}

		logger.Info("project found", "projectID", projectIDMaybe)

		data := Data{
			BackURL:   projectRoutes.ToProjectListPage(),
			Project:   project.EnsureComputed(),
			EditURL:   projectRoutes.ToProjectEdit(projectIDNew),
			DeleteURL: projectRoutes.ToProjectDelete(projectIDNew),
		}

		logger.Info("rendering project page")
		page.Respond(static.GetSiblingPath("projectPage.html"), data)(w, r)
	}
}
