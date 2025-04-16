package getProject

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/ctx/reqCtx"
	"imageresizerservice/app/error/errorPage"
	"imageresizerservice/app/projects/project"
	"imageresizerservice/app/projects/project/projectID"
	"imageresizerservice/app/projects/projectRoutes"
	"imageresizerservice/app/ui/page"
	"imageresizerservice/library/static"
	"net/http"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	mux.HandleFunc(projectRoutes.ProjectPage, Respond(ac))
}

type Data struct {
	BackURL   string
	Project   *project.Project
	EditURL   string
	DeleteURL string
}

func Respond(ac *appCtx.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := reqCtx.FromHttpRequest(ac, r)
		logger := req.Logger

		logger.Info("projectPage", "projectID", r.URL.Query().Get("projectID"))

		projectIDMaybe := r.URL.Query().Get("projectID")

		logger.Info("projectIDMaybe", "projectIDMaybe", projectIDMaybe)

		if projectIDMaybe == "" {
			logger.Error("missing project ID", "error", "Project ID is required")
			errorPage.Redirect(w, r, "Project ID is required")
			return
		}

		logger.Info("projectIDMaybe", "projectIDMaybe", projectIDMaybe)

		projectIDNew, err := projectID.New(projectIDMaybe)

		if err != nil {
			logger.Error("invalid project ID", "error", err)
			errorPage.Redirect(w, r, "Invalid project ID")
			return
		}

		uow, err := ac.UowFactory.Begin()
		if err != nil {
			logger.Error("database access failed", "error", err)
			errorPage.Redirect(w, r, "Failed to access database")
			return
		}
		defer uow.Rollback()

		project, err := ac.ProjectDB.GetByID(projectIDNew)

		if err != nil {
			logger.Error("project not found", "projectID", projectIDMaybe, "error", err)
			page.Respond(static.GetSiblingPath("notFound.html"), Data{
				BackURL: projectRoutes.ToProjectListPage(),
			})(w, r)
			return
		}

		if project == nil {
			logger.Error("project not found", "projectID", projectIDMaybe)
			page.Respond(static.GetSiblingPath("notFound.html"), Data{
				BackURL: projectRoutes.ToProjectListPage(),
			})(w, r)
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
		page.Respond(static.GetSiblingPath("page.html"), data)(w, r)
	}
}
