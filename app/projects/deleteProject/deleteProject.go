package deleteProject

import (
	"errors"
	"imageService/app/ctx/appCtx"
	"imageService/app/ctx/reqCtx"
	"imageService/app/home/homeRoutes"
	"imageService/app/projects/project"
	"imageService/app/projects/project/projectID"
	"imageService/app/projects/projectRoutes"
	"imageService/app/ui/breadcrumbs"
	"imageService/app/ui/confirmationPage"
	"imageService/app/ui/errorPage"
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

	confirmationPage.ConfirmationPage{
		Headline:    "Delete Project",
		Body:        "Are you sure you want to delete project '" + project.EnsureComputed().Name.String() + "'?",
		ConfirmURL:  r.URL.Path,
		ConfirmText: "Delete Project",
		CancelURL:   projectRoutes.ListProjects,
		CancelText:  "Cancel",
		HiddenForm: map[string]string{
			"projectID": projectIDVar.String(),
		},
		Breadcrumbs: []breadcrumbs.Breadcrumb{
			{Label: "Home", Href: homeRoutes.HomePage},
			{Label: "Projects", Href: projectRoutes.ListProjects},
			{Label: project.Name.String(), Href: project.EnsureComputed().URL},
			{Label: "Delete"},
		},
	}.Render(w, r)
}

func respondPost(ac *appCtx.AppCtx, w http.ResponseWriter, r *http.Request) {
	rc := reqCtx.FromHttpRequest(ac, r)
	logger := rc.Logger

	logger.Info("handling project delete request")

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

	err = deleteProject(ac, &rc, projectIDVar)
	if err != nil {
		logger.Error("failed to delete project", "error", err)
		errorPage.New(err).Redirect(w, r)
		return
	}

	logger.Info("project deleted successfully", "projectID", projectIDVar)
	http.Redirect(w, r, projectRoutes.ToListProjects(), http.StatusSeeOther)
}

func deleteProject(ac *appCtx.AppCtx, rc *reqCtx.ReqCtx, projectIDVar projectID.ProjectID) error {
	logger := rc.Logger

	uow, err := ac.UowFactory.Begin()
	if err != nil {
		logger.Error("failed to begin transaction", "error", err)
		return errors.New("failed to delete project")
	}
	defer uow.Rollback()

	existingProject, err := ac.ProjectDB.GetByID(projectIDVar)
	if err != nil {
		logger.Error("project not found", "projectID", projectIDVar, "error", err)
		return errors.New("project not found")
	}

	if existingProject == nil {
		logger.Error("project not found", "projectID", projectIDVar)
		return errors.New("project not found")
	}

	logger.Info("deleting project", "projectID", projectIDVar)

	if err = ac.ProjectDB.ZapByID(uow, projectIDVar); err != nil {
		logger.Error("failed to delete project", "error", err)
		return errors.New("failed to delete project")
	}

	if err = uow.Commit(); err != nil {
		logger.Error("failed to commit transaction", "error", err)
		return errors.New("failed to delete project")
	}

	return nil
}
