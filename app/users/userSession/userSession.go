package userSession

import (
	"imageresizerservice/app/users/userID"
	"imageresizerservice/app/users/userSession/userSessionID"
	"imageresizerservice/library/sessionID"
	"time"
)

type UserSession struct {
	ID        userSessionID.UserSessionID
	UserID    userID.UserID
	SessionID sessionID.SessionID
	CreatedAt time.Time
	EndedAt   time.Time
}
