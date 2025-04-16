package editProject

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/ctx/reqCtx"
	"imageresizerservice/app/projects/project"
	"imageresizerservice/app/projects/project/projectID"
	"imageresizerservice/app/projects/project/projectName"
	"imageresizerservice/app/projects/projectRoutes"
	"imageresizerservice/app/ui/page"
	"imageresizerservice/library/static"
	"net/http"
	"time"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	mux.HandleFunc(projectRoutes.ProjectEdit, Respond(ac))
}

type Data struct {
	Project     *project.Project
	ProjectPage string
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
		http.Error(w, "Project ID is required", http.StatusBadRequest)
		return
	}

	projectIDVar, err := projectID.New(projectIDMaybe)
	if err != nil {
		logger.Error("invalid project ID", "error", err)
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	uow, err := ac.UowFactory.Begin()
	if err != nil {
		logger.Error("database access failed", "error", err)
		http.Error(w, "Failed to access database", http.StatusInternalServerError)
		return
	}
	defer uow.Rollback()

	project, err := ac.ProjectDB.GetByID(projectIDVar)
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

	data := Data{
		Project:     project,
		ProjectPage: projectRoutes.ToProjectPage(projectIDVar),
	}

	page.Respond(static.GetSiblingPath("page.html"), data)(w, r)
}

func respondPost(ac *appCtx.AppCtx, w http.ResponseWriter, r *http.Request) {
	req := reqCtx.FromHttpRequest(ac, r)
	logger := req.Logger

	logger.Info("handling project edit request")

	// Handle form submission
	if err := r.ParseForm(); err != nil {
		logger.Error("failed to parse form", "error", err)
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	projectIDMaybe := r.FormValue("projectID")
	if projectIDMaybe == "" {
		logger.Error("missing project ID")
		http.Error(w, "Project ID is required", http.StatusBadRequest)
		return
	}

	projectIDVar, err := projectID.New(projectIDMaybe)
	if err != nil {
		logger.Error("invalid project ID", "error", err)
		http.Error(w, "Invalid project ID", http.StatusBadRequest)
		return
	}

	// Get existing project
	uow, err := ac.UowFactory.Begin()
	if err != nil {
		logger.Error("failed to begin transaction", "error", err)
		http.Error(w, "Failed to update project", http.StatusInternalServerError)
		return
	}
	defer uow.Rollback()

	existingProject, err := ac.ProjectDB.GetByID(projectIDVar)
	if err != nil {
		logger.Error("project not found", "projectID", projectIDMaybe, "error", err)
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}

	if existingProject == nil {
		logger.Error("project not found", "projectID", projectIDMaybe)
		http.Error(w, "Project not found", http.StatusNotFound)
		return
	}

	projectNameMaybe := r.FormValue("projectName")
	logger.Info("received project name", "projectName", projectNameMaybe)

	if projectNameMaybe == "" {
		logger.Error("empty project name")
		http.Error(w, "Project name is required", http.StatusBadRequest)
		return
	}

	projectNameVar, err := projectName.New(projectNameMaybe)
	if err != nil {
		logger.Error("invalid project name", "error", err)
		http.Error(w, "Invalid project name", http.StatusBadRequest)
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
		http.Error(w, "Failed to update project", http.StatusInternalServerError)
		return
	}

	if err = uow.Commit(); err != nil {
		logger.Error("failed to commit transaction", "error", err)
		http.Error(w, "Failed to update project", http.StatusInternalServerError)
		return
	}

	logger.Info("project updated successfully", "projectID", projectIDVar)
	http.Redirect(w, r, projectRoutes.ToProjectPage(projectIDVar), http.StatusSeeOther)
}
