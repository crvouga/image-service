package link

import (
	"time"

	"imageresizerservice/app/users/login/link/linkID"
	"imageresizerservice/library/email/emailAddress"
)

type Link struct {
	ID           linkID.LinkID
	EmailAddress emailAddress.EmailAddress
	CreatedAt    time.Time
	UsedAt       time.Time
}

func New(emailAddress emailAddress.EmailAddress) Link {
	return Link{
		ID:           linkID.Gen(),
		EmailAddress: emailAddress,
		CreatedAt:    time.Now(),
	}
}

func MarkAsUsed(l Link) Link {
	return Link{
		ID:           l.ID,
		EmailAddress: l.EmailAddress,
		CreatedAt:    l.CreatedAt,
		UsedAt:       time.Now(),
	}
}

func WasUsed(l *Link) bool {
	return !l.UsedAt.IsZero()
}
