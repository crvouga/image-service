package userSessionDb

import "imageresizerservice/users/userSession"

type UserSessionDb interface {
	GetById(id string) (*userSession.UserSession, error)
	Upsert(userSession.UserSession) error
}
