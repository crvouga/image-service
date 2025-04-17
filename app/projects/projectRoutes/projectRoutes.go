package projectRoutes

import (
	"fmt"
	"imageresizerservice/app/projects/project/projectID"
)

const (
	CreateProject = "/projects/create"
	EditProject   = "/projects/edit"
	DeleteProject = "/projects/delete"
	ListProjects  = "/projects/list"
	Project       = "/projects"
)

func withProjectID(route string, projectID projectID.ProjectID) string {
	return fmt.Sprintf("%s?projectID=%s", route, projectID)
}

func ToListProjects() string {
	return ListProjects
}

func ToEditProject(projectID projectID.ProjectID) string {
	return withProjectID(EditProject, projectID)
}

func ToGetProject(projectID projectID.ProjectID) string {
	return withProjectID(Project, projectID)
}

func ToCreateProject() string {
	return CreateProject
}

func ToDeleteProject(projectID projectID.ProjectID) string {
	return withProjectID(DeleteProject, projectID)
}
