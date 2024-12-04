package user_session_db

import (
	"image-resizer-service/login/user_session"
)

type UserSessionDb interface {
	GetById(id string) (*user_session.UserSession, error)
	Upsert(user_session.UserSession) error
}
