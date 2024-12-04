package login_link

import (
	"image-resizer-service/id"
	"time"
)

type LoginLink struct {
	Id        string
	Email     string
	CreatedAt time.Time
	UsedAt    time.Time
}

func New(email string) LoginLink {
	return LoginLink{
		Id:        id.Gen(),
		Email:     email,
		CreatedAt: time.Now(),
	}
}

func MarkAsUsed(l LoginLink) LoginLink {
	l.UsedAt = time.Now()
	return l
}

func WasUsed(l *LoginLink) bool {

	return !l.UsedAt.IsZero()
}
