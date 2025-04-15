package userSessionDb

import (
	"imageresizerservice/app/ctx/sessionID"
	"imageresizerservice/app/users/userSession"
	"imageresizerservice/library/uow"
)

type UserSessionDb interface {
	GetById(id sessionID.SessionID) (*userSession.UserSession, error)
	Upsert(uow *uow.Uow, userSession userSession.UserSession) error
}
