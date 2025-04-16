package projectListPage

import (
	"imageresizerservice/app/ctx/appCtx"
	"imageresizerservice/app/ctx/reqCtx"
	"imageresizerservice/app/home/homeRoutes"
	"imageresizerservice/app/projects/project"
	"imageresizerservice/app/projects/projectRoutes"
	"imageresizerservice/app/ui/page"
	"imageresizerservice/app/users/userAccount/userAccountRoutes"
	"imageresizerservice/library/static"
	"net/http"
)

func Router(mux *http.ServeMux, appCtx *appCtx.AppCtx) {
	mux.HandleFunc(projectRoutes.ProjectListPage, Respond(appCtx))
}

type Data struct {
	BackURL             string
	UserAccountPageHref string
	Projects            []*project.Project
	CreateURL           string
}

func Respond(appCtx *appCtx.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := reqCtx.FromHttpRequest(appCtx, r)
		logger := req.Logger
		createdByUserID := req.UserSession.UserID

		logger.Info("projectListPage", "userID", createdByUserID)

		uow, err := appCtx.UowFactory.Begin()
		if err != nil {
			logger.Error("database access failed", "error", err)
			http.Error(w, "Failed to access database", http.StatusInternalServerError)
			return
		}
		defer uow.Rollback()

		projects, err := appCtx.ProjectDB.GetByCreatedByUserID(createdByUserID)
		if err != nil {
			logger.Error("failed to fetch projects", "userID", createdByUserID, "error", err)
			http.Error(w, "Failed to fetch projects", http.StatusInternalServerError)
			return
		}

		for _, project := range projects {
			project.EnsureComputed()
		}

		data := Data{
			BackURL:             homeRoutes.HomePage,
			UserAccountPageHref: userAccountRoutes.UserAccountPage,
			Projects:            projects,
			CreateURL:           projectRoutes.ToProjectCreate(),
		}

		logger.Info("rendering project list page", "projectCount", len(projects))
		page.Respond(static.GetSiblingPath("projectListPage.html"), data)(w, r)
	}
}
