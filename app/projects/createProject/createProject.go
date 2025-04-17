package createProject

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
	mux.HandleFunc(projectRoutes.CreateProject, Respond(ac))
}

type Data struct {
	HomeURL     string
	ProjectsURL string
}

func Respond(ac *appCtx.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			respondPost(ac, w, r)
		} else {
			respondGet(w, r)
		}
	}
}

func respondGet(w http.ResponseWriter, r *http.Request) {
	data := Data{
		HomeURL:     homeRoutes.HomePage,
		ProjectsURL: projectRoutes.ListProjects,
	}
	page.Respond(data, static.GetSiblingPath("createProject.html"))(w, r)
}
func respondPost(ac *appCtx.AppCtx, w http.ResponseWriter, r *http.Request) {
	req := reqCtx.FromHttpRequest(ac, r)
	logger := req.Logger

	logger.Info("handling project creation request")

	// Handle form submission
	if err := r.ParseForm(); err != nil {
		logger.Error("failed to parse form", "error", err)
		errorPage.New(errors.New("failed to parse form")).Redirect(w, r)
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

	projectID := projectID.Gen()

	projectNew := project.Project{
		ID:              projectID,
		CreatedByUserID: req.UserSession.UserID,
		Name:            projectNameVar,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		AllowedDomains:  allowedDomainsList,
	}

	logger.Info("projectNew", "projectNew", projectNew)

	logger.Info("creating new project", "projectID", projectID, "createdByUserID", projectNew.CreatedByUserID)

	uow, err := ac.UowFactory.Begin()

	if err != nil {
		logger.Error("failed to begin transaction", "error", err)
		errorPage.New(errors.New("failed to create project")).Redirect(w, r)
		return
	}

	logger.Info("upserting project", "projectID", projectID, "createdByUserID", projectNew.CreatedByUserID)

	if err = ac.ProjectDB.Upsert(uow, &projectNew); err != nil {
		logger.Error("failed to upsert project", "error", err)
		errorPage.New(errors.New("failed to create project")).Redirect(w, r)
		return
	}

	if err = uow.Commit(); err != nil {
		logger.Error("failed to commit transaction", "error", err)
		errorPage.New(errors.New("failed to create project")).Redirect(w, r)
		return
	}

	logger.Info("project created successfully", "projectID", projectID)
	http.Redirect(w, r, projectRoutes.ToGetProject(projectID), http.StatusSeeOther)
}
