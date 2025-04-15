package userSessionDb

import (
	"imageresizerservice/app/users/userSession"
	"imageresizerservice/library/uow"
)

type UserSessionDb interface {
	GetById(id string) (*userSession.UserSession, error)
	Upsert(uow *uow.Uow, userSession userSession.UserSession) error
}
