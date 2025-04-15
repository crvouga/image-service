package link

import (
	"time"

	"imageresizerservice/library/email/emailAddress"
	"imageresizerservice/library/id"
)

type Link struct {
	Id           string
	EmailAddress emailAddress.EmailAddress
	CreatedAt    time.Time
	UsedAt       time.Time
}

func New(emailAddress emailAddress.EmailAddress) Link {
	return Link{
		Id:           id.Gen(),
		EmailAddress: emailAddress,
		CreatedAt:    time.Now(),
	}
}

func MarkAsUsed(l Link) Link {
	return Link{
		Id:           l.Id,
		EmailAddress: l.EmailAddress,
		CreatedAt:    l.CreatedAt,
		UsedAt:       time.Now(),
	}
}

func WasUsed(l *Link) bool {
	return !l.UsedAt.IsZero()
}
