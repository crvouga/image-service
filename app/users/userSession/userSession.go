package userSession

import "time"

type UserSession struct {
	ID        string
	UserID    string
	SessionID string
	CreatedAt time.Time
	EndedAt   time.Time
}
