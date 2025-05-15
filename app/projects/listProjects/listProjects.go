package listProjects

import (
	"imageService/app/ctx/appCtx"
	"imageService/app/ctx/reqCtx"
	"imageService/app/home/homeRoutes"
	"imageService/app/projects/project"
	"imageService/app/projects/projectRoutes"
	"imageService/app/ui/breadcrumbs"
	"imageService/app/ui/errorPage"
	"imageService/app/ui/page"
	"imageService/app/ui/pageHeader"
	"imageService/library/static"
	"net/http"
)

func Router(mux *http.ServeMux, ac *appCtx.AppCtx) {
	mux.HandleFunc(projectRoutes.ListProjects, Respond(ac))
}

type Data struct {
	Projects    []*project.Project
	Breadcrumbs breadcrumbs.Breadcrumbs
	PageHeader  pageHeader.PageHeader
}

func Respond(ac *appCtx.AppCtx) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := reqCtx.FromHttpRequest(ac, r)
		logger := req.Logger
		createdByUserID := req.UserSession.UserID

		logger.Info("projectListPage", "userID", createdByUserID)

		uow, err := ac.UowFactory.Begin()
		if err != nil {
			logger.Error("failed to fetch projects", "userID", createdByUserID, "error", err)
			errorPage.New(err).Redirect(w, r)
			return
		}
		defer uow.Rollback()

		projects, err := ac.ProjectDB.GetByCreatedByUserID(createdByUserID)
		if err != nil {
			logger.Error("failed to fetch projects", "userID", createdByUserID, "error", err)
			errorPage.New(err).Redirect(w, r)
			return
		}

		for _, project := range projects {
			project.EnsureComputed()
		}

		data := Data{
			Projects: projects,
			Breadcrumbs: []breadcrumbs.Breadcrumb{
				{Label: "Home", Href: homeRoutes.HomePage},
				{Label: "Projects"},
			},
			PageHeader: pageHeader.PageHeader{
				Title: "Projects",
				Actions: []pageHeader.Action{
					{
						Label: "Create",
						URL:   projectRoutes.ToCreateProject(),
					},
				},
			},
		}

		logger.Info("rendering project list page", "projectCount", len(projects))
		page.Respond(data, static.GetSiblingPath("listProjects.html"))(w, r)
	}
}
