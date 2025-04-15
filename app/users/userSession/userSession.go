package userSession

import "time"

type UserSession struct {
	Id        string
	UserId    string
	SessionId string
	CreatedAt time.Time
	EndedAt   time.Time
}
