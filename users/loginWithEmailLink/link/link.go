package link

import (
	"time"

	"imageresizerservice/id"
)

type Link struct {
	Id           string
	EmailAddress string
	CreatedAt    time.Time
	UsedAt       time.Time
}

func New(email string) Link {
	return Link{
		Id:           id.Gen(),
		EmailAddress: email,
		CreatedAt:    time.Now(),
	}
}

func MarkAsUsed(l Link) Link {
	l.UsedAt = time.Now()
	return l
}

func WasUsed(l *Link) bool {
	return !l.UsedAt.IsZero()
}
