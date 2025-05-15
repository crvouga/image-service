package userSession

import (
	"imageService/app/users/userID"
	"imageService/app/users/userSession/userSessionID"
	"imageService/library/sessionID"
	"time"
)

type UserSession struct {
	ID        userSessionID.UserSessionID
	UserID    userID.UserID
	SessionID sessionID.SessionID
	CreatedAt time.Time
	EndedAt   time.Time
}
