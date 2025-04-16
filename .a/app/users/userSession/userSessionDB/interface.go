package userSessionDB

import (
	"imageresizerservice/app/ctx/sessionID"
	"imageresizerservice/app/users/userSession"
	"imageresizerservice/library/uow"
)

type UserSessionDB interface {
	GetBySessionID(id sessionID.SessionID) (*userSession.UserSession, error)
	Upsert(uow *uow.Uow, userSession userSession.UserSession) error
	ZapBySessionID(uow *uow.Uow, sessionID sessionID.SessionID) error
}
