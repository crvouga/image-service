package loginLink

import (
	"time"

	"imageresizerservice/id"
)

type LoginLink struct {
	Id           string
	EmailAddress string
	CreatedAt    time.Time
	UsedAt       time.Time
}

func New(email string) LoginLink {
	return LoginLink{
		Id:           id.Gen(),
		EmailAddress: email,
		CreatedAt:    time.Now(),
	}
}

func MarkAsUsed(l LoginLink) LoginLink {
	l.UsedAt = time.Now()
	return l
}

func WasUsed(l *LoginLink) bool {
	return !l.UsedAt.IsZero()
}
