package userSession

import (
	"imageresizerservice/app/ctx/sessionID"
	"imageresizerservice/app/users/userID"
	"imageresizerservice/app/users/userSession/userSessionID"
	"time"
)

type UserSession struct {
	ID        userSessionID.UserSessionID
	UserID    userID.UserID
	SessionID sessionID.SessionID
	CreatedAt time.Time
	EndedAt   time.Time
}
