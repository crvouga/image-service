package projectDB

import (
	"imageService/app/projects/project"
	"imageService/app/projects/project/projectID"
	"imageService/app/users/userID"
	"imageService/library/uow"
)

type ProjectDB interface {
	GetByID(projectID projectID.ProjectID) (*project.Project, error)
	Upsert(uow *uow.Uow, project *project.Project) error
	GetByCreatedByUserID(createdByUserID userID.UserID) ([]*project.Project, error)
	ZapByID(uow *uow.Uow, projectID projectID.ProjectID) error
}
