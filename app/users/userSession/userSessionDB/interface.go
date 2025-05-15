package userSessionDB

import (
	"imageService/app/users/userSession"
	"imageService/library/sessionID"
	"imageService/library/uow"
)

type UserSessionDB interface {
	GetBySessionID(id sessionID.SessionID) (*userSession.UserSession, error)
	Upsert(uow *uow.Uow, userSession userSession.UserSession) error
	ZapBySessionID(uow *uow.Uow, sessionID sessionID.SessionID) error
}
