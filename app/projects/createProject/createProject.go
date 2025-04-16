package createProject

import (
	"imageresizerservice/app/ctx/appContext"
	"imageresizerservice/app/ctx/reqCtx"
	"imageresizerservice/app/home/homeRoutes"
	"imageresizerservice/app/projects/project"
	"imageresizerservice/app/projects/project/projectID"
	"imageresizerservice/app/projects/project/projectName"
	"imageresizerservice/app/projects/projectRoutes"
	"imageresizerservice/app/ui/page"
	"imageresizerservice/library/static"
	"net/http"
	"time"
)

func Router(mux *http.ServeMux, ac *appContext.AppCtx) {
	mux.HandleFunc(projectRoutes.ProjectCreate, Respond(ac))
}

type Data struct {
	HomePage string
}

func Respond(ac *appContext.AppCtx) http.HandlerFunc {
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
		HomePage: homeRoutes.HomePage,
	}
	page.Respond(static.GetSiblingPath("page.html"), data)(w, r)
}
func respondPost(ac *appContext.AppCtx, w http.ResponseWriter, r *http.Request) {
	req := reqCtx.FromHttpRequest(ac, r)
	logger := req.Logger

	logger.Info("handling project creation request")

	// Handle form submission
	if err := r.ParseForm(); err != nil {
		logger.Error("failed to parse form", "error", err)
		http.Error(w, "Failed to parse form", http.StatusBadRequest)
		return
	}

	projectNameMaybe := r.FormValue("projectName")
	logger.Info("received project name", "projectName", projectNameMaybe)

	if projectNameMaybe == "" {
		logger.Error("empty project name")
		http.Error(w, "Project name is required", http.StatusBadRequest)
		return
	}

	projectNameInst, err := projectName.New(projectNameMaybe)

	if err != nil {
		logger.Error("invalid project name", "error", err)
		http.Error(w, "Invalid project name", http.StatusBadRequest)
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
		Name:            projectNameInst,
		CreatedAt:       time.Now(),
		UpdatedAt:       time.Now(),
		AllowedDomains:  allowedDomainsList,
	}

	logger.Info("projectNew", "projectNew", projectNew)

	logger.Info("creating new project", "projectID", projectID, "createdByUserID", projectNew.CreatedByUserID)

	uow, err := ac.UowFactory.Begin()

	if err != nil {
		logger.Error("failed to begin transaction", "error", err)
		http.Error(w, "Failed to create project", http.StatusInternalServerError)
		return
	}

	logger.Info("upserting project", "projectID", projectID, "createdByUserID", projectNew.CreatedByUserID)

	if err = ac.ProjectDB.Upsert(uow, &projectNew); err != nil {
		logger.Error("failed to upsert project", "error", err)
		http.Error(w, "Failed to create project", http.StatusInternalServerError)
		return
	}

	if err = uow.Commit(); err != nil {
		logger.Error("failed to commit transaction", "error", err)
		http.Error(w, "Failed to create project", http.StatusInternalServerError)
		return
	}

	logger.Info("project created successfully", "projectID", projectID)
	http.Redirect(w, r, projectRoutes.ToProjectPage(projectID), http.StatusSeeOther)
}
