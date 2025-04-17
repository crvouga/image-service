package projectPage

import (
	"errors"
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/ctx/reqCtx"
	"imageresizerservice/app/home/homeRoutes"
	"imageresizerservice/app/projects/project"
	"imageresizerservice/app/projects/project/projectID"
	"imageresizerservice/app/projects/projectRoutes"
	"imageresizerservice/app/ui/breadcrumbs"
	"imageresizerservice/app/ui/errorPage"
	"imageresizerservice/app/ui/notFoundPage"
	"imageresizerservice/app/ui/page"
	"imageresizerservice/app/ui/pageHeader"
	"imageresizerservice/library/static"
	"net/http"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	mux.HandleFunc(projectRoutes.Project, Respond(ac))
}

type Data struct {
	Project     *project.Project
	Breadcrumbs breadcrumbs.Breadcrumbs
	PageHeader  pageHeader.PageHeader
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
			errorPage.New(errors.New("project ID is required")).Redirect(w, r)
			return
		}

		logger.Info("projectIDMaybe", "projectIDMaybe", projectIDMaybe)

		projectIDNew, err := projectID.New(projectIDMaybe)

		if err != nil {
			logger.Error("invalid project ID", "error", err)
			errorPage.New(errors.New("invalid project ID")).Redirect(w, r)
			return
		}

		uow, err := ac.UowFactory.Begin()
		if err != nil {
			logger.Error("database access failed", "error", err)
			errorPage.New(errors.New("failed to access database")).Redirect(w, r)
			return
		}
		defer uow.Rollback()

		project, err := ac.ProjectDB.GetByID(projectIDNew)

		if err != nil {
			logger.Error("project not found", "projectID", projectIDMaybe, "error", err)
			notFoundPage.New(projectRoutes.ListProjects, "Back to Projects").Redirect(w, r)
			return
		}

		if project == nil {
			notFoundPage.New(projectRoutes.ListProjects, "Back to Projects").Redirect(w, r)
			return
		}

		logger.Info("project found", "projectID", projectIDMaybe)

		data := Data{
			Project: project.EnsureComputed(),
			Breadcrumbs: []breadcrumbs.Breadcrumb{
				{Label: "Home", Href: homeRoutes.HomePage},
				{Label: "Projects", Href: projectRoutes.ListProjects},
				{Label: project.EnsureComputed().Name.String()},
			},
			PageHeader: pageHeader.PageHeader{
				Title: project.EnsureComputed().Name.String(),
				Actions: []pageHeader.Action{
					{
						Label: "Edit",
						URL:   project.EnsureComputed().EditURL,
					},
					{
						Label: "Delete",
						URL:   project.EnsureComputed().DeleteURL,
					},
				},
			},
		}

		logger.Info("rendering project page")
		page.Respond(data, static.GetSiblingPath("projectPage.html"))(w, r)
	}
}
