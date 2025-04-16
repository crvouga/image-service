package projectDB

import (
	"imageresizerservice/app/projects/project"
	"imageresizerservice/app/projects/project/projectID"
	"imageresizerservice/library/uow"
)

type ProjectDB interface {
	GetByID(id projectID.ProjectID) (*project.Project, error)
	Upsert(uow *uow.Uow, project project.Project) error
	ZapByID(uow *uow.Uow, id projectID.ProjectID) error
}
