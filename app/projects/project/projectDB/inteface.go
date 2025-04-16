package projectDB

import (
	"imageresizerservice/app/projects/project"
	"imageresizerservice/app/projects/project/projectID"
	"imageresizerservice/app/users/userID"
	"imageresizerservice/library/uow"
)

type ProjectDB interface {
	GetByID(projectID projectID.ProjectID) (*project.Project, error)
	Upsert(uow *uow.Uow, project *project.Project) error
	GetByCreatedByUserID(createdByUserID userID.UserID) ([]*project.Project, error)
	ZapByID(uow *uow.Uow, projectID projectID.ProjectID) error
}
