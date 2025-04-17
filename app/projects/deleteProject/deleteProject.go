package deleteProject

import (
	"errors"
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/ctx/reqCtx"
	"imageresizerservice/app/home/homeRoutes"
	"imageresizerservice/app/projects/project"
	"imageresizerservice/app/projects/project/projectID"
	"imageresizerservice/app/projects/projectRoutes"
	"imageresizerservice/app/ui/errorPage"
	"imageresizerservice/app/ui/page"
	"imageresizerservice/library/static"
	"net/http"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	mux.HandleFunc(projectRoutes.DeleteProject, Respond(ac))
}

type Data struct {
	HomeURL     string
	ProjectsURL string
	Project     *project.Project
}

func Respond(ac *appCtx.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			respondPost(ac, w, r)
		} else {
			respondGet(ac, w, r)
		}
	}
}

func respondGet(ac *appCtx.AppCtx, w http.ResponseWriter, r *http.Request) {
	req := reqCtx.FromHttpRequest(ac, r)
	logger := req.Logger

	projectIDMaybe := r.URL.Query().Get("projectID")
	if projectIDMaybe == "" {
		logger.Error("missing project ID")
		errorPage.New(errors.New("project id is required")).Redirect(w, r)
		return
	}

	projectIDVar, err := projectID.New(projectIDMaybe)
	if err != nil {
		logger.Error("invalid project ID", "error", err)
		errorPage.New(errors.New("invalid project id")).Redirect(w, r)
		return
	}

	uow, err := ac.UowFactory.Begin()
	if err != nil {
		logger.Error("database access failed", "error", err)
		errorPage.New(errors.New("failed to access database")).Redirect(w, r)
		return
	}
	defer uow.Rollback()

	project, err := ac.ProjectDB.GetByID(projectIDVar)
	if err != nil {
		logger.Error("project not found", "projectID", projectIDMaybe, "error", err)
		errorPage.New(errors.New("project not found")).Redirect(w, r)
		return
	}

	if project == nil {
		logger.Error("project not found", "projectID", projectIDMaybe)
		errorPage.New(errors.New("project not found")).Redirect(w, r)
		return
	}

	data := Data{
		Project:     project.EnsureComputed(),
		HomeURL:     homeRoutes.HomePage,
		ProjectsURL: projectRoutes.ListProjects,
	}

	page.Respond(data, static.GetSiblingPath("deleteProject.html"))(w, r)
}

func respondPost(ac *appCtx.AppCtx, w http.ResponseWriter, r *http.Request) {
	req := reqCtx.FromHttpRequest(ac, r)
	logger := req.Logger

	logger.Info("handling project delete request")

	// Handle form submission
	if err := r.ParseForm(); err != nil {
		logger.Error("failed to parse form", "error", err)
		errorPage.New(errors.New("failed to parse form")).Redirect(w, r)
		return
	}

	projectIDMaybe := r.FormValue("projectID")
	if projectIDMaybe == "" {
		logger.Error("missing project ID")
		errorPage.New(errors.New("project id is required")).Redirect(w, r)
		return
	}

	projectIDVar, err := projectID.New(projectIDMaybe)
	if err != nil {
		logger.Error("invalid project ID", "error", err)
		errorPage.New(errors.New("invalid project id")).Redirect(w, r)
		return
	}

	// Get existing project
	uow, err := ac.UowFactory.Begin()
	if err != nil {
		logger.Error("failed to begin transaction", "error", err)
		errorPage.New(errors.New("failed to delete project")).Redirect(w, r)
		return
	}
	defer uow.Rollback()

	existingProject, err := ac.ProjectDB.GetByID(projectIDVar)
	if err != nil {
		logger.Error("project not found", "projectID", projectIDMaybe, "error", err)
		errorPage.New(errors.New("project not found")).Redirect(w, r)
		return
	}

	if existingProject == nil {
		logger.Error("project not found", "projectID", projectIDMaybe)
		errorPage.New(errors.New("project not found")).Redirect(w, r)
		return
	}

	logger.Info("deleting project", "projectID", projectIDVar)

	if err = ac.ProjectDB.ZapByID(uow, projectIDVar); err != nil {
		logger.Error("failed to delete project", "error", err)
		errorPage.New(errors.New("failed to delete project")).Redirect(w, r)
		return
	}

	if err = uow.Commit(); err != nil {
		logger.Error("failed to commit transaction", "error", err)
		errorPage.New(errors.New("failed to delete project")).Redirect(w, r)
		return
	}

	logger.Info("project deleted successfully", "projectID", projectIDVar)
	http.Redirect(w, r, projectRoutes.ToListProjects(), http.StatusSeeOther)
}
