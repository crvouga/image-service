package editProject

import (
	"errors"
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/ctx/reqCtx"
	"imageresizerservice/app/home/homeRoutes"
	"imageresizerservice/app/projects/project"
	"imageresizerservice/app/projects/project/projectID"
	"imageresizerservice/app/projects/project/projectName"
	"imageresizerservice/app/projects/projectRoutes"
	"imageresizerservice/app/ui/errorPage"
	"imageresizerservice/app/ui/page"
	"imageresizerservice/library/static"
	"net/http"
	"time"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	mux.HandleFunc(projectRoutes.EditProject, Respond(ac))
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

type Data struct {
	Project     *project.Project
	ProjectsURL string
	HomeURL     string
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
		ProjectsURL: projectRoutes.ToListProjects(),
		HomeURL:     homeRoutes.HomePage,
	}

	page.Respond(data, static.GetSiblingPath("editProject.html"))(w, r)
}

func respondPost(ac *appCtx.AppCtx, w http.ResponseWriter, r *http.Request) {
	req := reqCtx.FromHttpRequest(ac, r)
	logger := req.Logger

	logger.Info("handling project edit request")

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
		errorPage.New(errors.New("failed to update project")).Redirect(w, r)
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

	projectNameMaybe := r.FormValue("projectName")
	logger.Info("received project name", "projectName", projectNameMaybe)

	if projectNameMaybe == "" {
		logger.Error("empty project name")
		errorPage.New(errors.New("project name is required")).Redirect(w, r)
		return
	}

	projectNameVar, err := projectName.New(projectNameMaybe)
	if err != nil {
		logger.Error("invalid project name", "error", err)
		errorPage.New(errors.New("invalid project name")).Redirect(w, r)
		return
	}

	allowedDomainsLines := r.FormValue("allowedDomains")
	logger.Info("received allowed domains", "allowedDomains", allowedDomainsLines)

	allowedDomainsList := project.UrlLinesToUrlList(allowedDomainsLines)
	logger.Info("parsed allowed domains", "count", len(allowedDomainsList))

	// Update project with new values
	updatedProject := project.Project{
		ID:              existingProject.ID,
		CreatedByUserID: existingProject.CreatedByUserID,
		Name:            projectNameVar,
		CreatedAt:       existingProject.CreatedAt,
		UpdatedAt:       time.Now(),
		AllowedDomains:  allowedDomainsList,
	}

	logger.Info("updating project", "projectID", projectIDVar)

	if err = ac.ProjectDB.Upsert(uow, &updatedProject); err != nil {
		logger.Error("failed to update project", "error", err)
		errorPage.New(errors.New("failed to update project")).Redirect(w, r)
		return
	}

	if err = uow.Commit(); err != nil {
		logger.Error("failed to commit transaction", "error", err)
		errorPage.New(errors.New("failed to update project")).Redirect(w, r)
		return
	}

	logger.Info("project updated successfully", "projectID", projectIDVar)
	http.Redirect(w, r, projectRoutes.ToGetProject(projectIDVar), http.StatusSeeOther)
}
