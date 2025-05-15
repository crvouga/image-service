package apiDocsPage

import (
	"imageService/app/ctx/appCtx"
	"imageService/app/ctx/reqCtx"
	"imageService/app/projects/project"
	"imageService/app/projects/projectRoutes"
	"imageService/app/users/userID"
	"net/http"
)

type ProjectIDSelect struct {
	Projects         []*project.Project
	CreateProjectURL string
	ProjectID        string
}

func GetProjectIDSelect(ac *appCtx.AppCtx, r *http.Request) ProjectIDSelect {
	rc := reqCtx.FromHttpRequest(ac, r)

	return ProjectIDSelect{
		Projects:         getProjects(ac, rc.UserSession.UserID),
		ProjectID:        r.URL.Query().Get("projectID"),
		CreateProjectURL: projectRoutes.ToCreateProject(),
	}
}

func getProjects(ac *appCtx.AppCtx, userID userID.UserID) []*project.Project {
	projects, err := ac.ProjectDB.GetByCreatedByUserID(userID)

	if err != nil {
		return []*project.Project{}
	}

	return projects
}
