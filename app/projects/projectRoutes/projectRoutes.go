package projectRoutes

import (
	"fmt"
	"imageresizerservice/app/projects/project/projectID"
)

const (
	ProjectCreate = "/projects-create"
	ProjectEdit   = "/projects-edit"
	ProjectDelete = "/projects-delete"
	ProjectPage   = "/projects"
)

func withProjectID(route string, projectID projectID.ProjectID) string {
	return fmt.Sprintf("%s?projectID=%s", route, projectID)
}

func ToProjectEdit(projectID projectID.ProjectID) string {
	return withProjectID(ProjectEdit, projectID)
}

func ToProjectPage(projectID projectID.ProjectID) string {
	return withProjectID(ProjectPage, projectID)
}

func ToProjectCreate() string {
	return ProjectCreate
}

func ToProjectDelete(projectID projectID.ProjectID) string {
	return withProjectID(ProjectDelete, projectID)
}
