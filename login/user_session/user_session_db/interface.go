package user_session_db

import (
	"imageresizerservice.com/login/user_session"
)

type UserSessionDb interface {
	GetById(id string) (*user_session.UserSession, error)
	Upsert(user_session.UserSession) error
}
